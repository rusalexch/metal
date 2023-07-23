package config

import (
	"crypto/rsa"
	"crypto/x509"
	"flag"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	// переменная для адреса сервера.
	addr *string
	// переменная для интервала сбора метрик.
	reportInterval time.Duration
	// переменная для интервала отправки метрик на сервер.
	pollInterval time.Duration
	// интервал сохранения в файловое хранилище.
	storeInterval time.Duration
	// путь к файлу файлового хранилища.
	storeFile *string
	// флаг подгружать ли сохраненные данные из файлового хранилища, или начинать с чистого файла.
	restore *string
	// ключ хэш-функции.
	key *string
	// url строка подключения базы данных.
	dbURL *string
	// количество одновременно исходящих запросов от агента.
	rateLimit int
	// cryptoKeyPath ключ, для агента публичный для сервера приватный
	cryptoKeyPath *string
)

func init() {
	addr = flag.String("a", defaultAddr, "set address")
	pollInterval = defaultPoolInterval
	flag.Func("p", "poll interval", func(s string) (err error) {
		reportInterval, err = time.ParseDuration(s)
		if err != nil {
			return err
		}
		return nil
	})
	storeInterval = defaultStoreInterval
	flag.Func("i", "store interval", func(i string) (err error) {
		storeInterval, err = time.ParseDuration(i)
		if err != nil {
			return err
		}
		return nil
	})
	storeFile = flag.String("f", defaultStoreFile, "store file")
	key = flag.String("k", defaultKey, "hash secret key")
	dbURL = flag.String("d", "", "database url string")
	rateLimit = defaultRateLimit
	flag.Func("l", "rate limit", func(i string) (err error) {
		rateLimit, err = strconv.Atoi(i)
		if err != nil {
			return err
		}
		return nil
	})
	cryptoKeyPath = flag.String("crypto-key", "", "set crypto key file (public for agent, private for server)")
}

// NewAgentConfig - конструктор конфигурации для агента.
func NewAgentConfig() AgentConfig {
	reportInterval = defaultReportInterval
	flag.Func("r", "report interval", func(s string) (err error) {
		reportInterval, err = time.ParseDuration(s)
		if err != nil {
			return err
		}
		return nil
	})
	flag.Parse()
	parseENV()

	cryptoKey, err := getPublicKey(cryptoKeyPath)
	if err != nil {
		log.Println("can't get public crypto key")
		log.Fatal(err)
	}

	return AgentConfig{
		Addr:           *addr,
		ReportInterval: reportInterval,
		PoolInterval:   pollInterval,
		HashKey:        *key,
		RateLimit:      rateLimit,
		PublicKey:      cryptoKey,
	}
}

// NewServerConfig - конструктор конфигурации для сервера.
func NewServerConfig() ServerConfig {
	restore = flag.String("r", defaultRestore, "is restore from file")
	flag.Parse()
	parseENV()

	cryptoKey, err := getPrivateKey(cryptoKeyPath)
	if err != nil {
		log.Println("can't get private crypto key")
		log.Fatal(err)
	}

	return ServerConfig{
		Addr:          *addr,
		StoreInterval: storeInterval,
		StoreFile:     *storeFile,
		Restore:       *restore == "true",
		HashKey:       *key,
		DBURL:         *dbURL,
		PrivateKey:    cryptoKey,
	}
}

// parseENV - метод парсинга переменных окружения.
func parseENV() {
	if addrEnv, isSet := os.LookupEnv("ADDRESS"); isSet {
		addr = &addrEnv
	}
	if reportIntervalEnv, isSet := os.LookupEnv("REPORT_INTERVAL"); isSet {
		t, err := time.ParseDuration(reportIntervalEnv)
		if err != nil {
			log.Fatal(err)
		}
		reportInterval = t
	}
	if poolIntervalEnv, isSet := os.LookupEnv("POLL_INTERVAL"); isSet {
		t, err := time.ParseDuration(poolIntervalEnv)
		if err != nil {
			log.Fatal(err)
		}
		pollInterval = t
	}
	if storeIntervalEnv, isSet := os.LookupEnv("STORE_INTERVAL"); isSet {
		s, err := time.ParseDuration(storeIntervalEnv)
		if err != nil {
			log.Fatal(err)
		}
		storeInterval = s
	}
	if storeFileEnv, isSet := os.LookupEnv("STORE_FILE"); isSet {
		storeFile = &storeFileEnv
	}
	if restoreEnv, isSet := os.LookupEnv("RESTORE"); isSet {
		restore = &restoreEnv
	}
	if keyEnv, isSet := os.LookupEnv("KEY"); isSet {
		key = &keyEnv
	}
	if dbURLEnv, isSet := os.LookupEnv("DATABASE_DSN"); isSet {
		dbURL = &dbURLEnv
	}
	if rateLimitEnv, isSet := os.LookupEnv("RATE_LIMIT"); isSet {
		limit, err := strconv.Atoi(rateLimitEnv)
		if err != nil {
			log.Fatal(err)
		}
		rateLimit = limit
	}
	if cryptoKeyPathEnv, isSet := os.LookupEnv("CRYPTO_KEY"); isSet {
		cryptoKeyPath = &cryptoKeyPathEnv
	}
}

func getPublicKey(path *string) (*rsa.PublicKey, error) {
	if path == nil || *path == "" {
		return nil, nil
	}
	key, err := readCryptoFile(*path)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PublicKey(key)
}

func getPrivateKey(path *string) (*rsa.PrivateKey, error) {
	if path == nil || *path == "" {
		return nil, nil
	}
	key, err := readCryptoFile(*path)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(key)
}

func readCryptoFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	return file, nil
}
