package config

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/rusalexch/metal/internal/config/json"
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
	// defValues - параметры по умолчанию, или настройки json-файла
	defValues *json.DefaultConfig
)

func init() {
	defValues = json.ParseJson()

	addr = flag.String("a", defValues.Address, "set address")
	pollInterval = defValues.PollInterval
	flag.Func("p", "poll interval", func(s string) (err error) {
		reportInterval, err = time.ParseDuration(s)
		if err != nil {
			return err
		}
		return nil
	})
	storeInterval = defValues.StoreInterval
	flag.Func("i", "store interval", func(i string) (err error) {
		storeInterval, err = time.ParseDuration(i)
		if err != nil {
			return err
		}
		return nil
	})
	storeFile = flag.String("f", defValues.StoreFile, "store file")
	key = flag.String("k", defValues.Key, "hash secret key")
	dbURL = flag.String("d", defValues.DatabaseDSN, "database url string")
	rateLimit = defValues.RateLimit
	flag.Func("l", "rate limit", func(i string) (err error) {
		rateLimit, err = strconv.Atoi(i)
		if err != nil {
			return err
		}
		return nil
	})
	cryptoKeyPath = flag.String("crypto-key", defValues.CryptoKey, "set crypto key file (public for agent, private for server)")

}

// NewAgentConfig - конструктор конфигурации для агента.
func NewAgentConfig() AgentConfig {
	reportInterval = defValues.ReportInterval
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
	restore = flag.String("r", defValues.Restore, "is restore from file")
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
