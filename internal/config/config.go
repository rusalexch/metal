package config

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/rusalexch/metal/internal/config/json"
	"golang.org/x/exp/constraints"
)

var (
	// переменная для адреса сервера.
	addr *string
	// переменная для интервала сбора метрик.
	reportInterval *time.Duration
	// переменная для интервала отправки метрик на сервер.
	pollInterval *time.Duration
	// интервал сохранения в файловое хранилище.
	storeInterval *time.Duration
	// путь к файлу файлового хранилища.
	storeFile *string
	// флаг подгружать ли сохраненные данные из файлового хранилища, или начинать с чистого файла.
	restore *string
	// ключ хэш-функции.
	key *string
	// url строка подключения базы данных.
	dbURL *string
	// количество одновременно исходящих запросов от агента.
	rateLimit *int
	// cryptoKeyPath ключ, для агента публичный для сервера приватный
	cryptoKeyPath *string
	// jsonFile - путь к файлу конфигурации json
	jsonFile string
)

func init() {
	flag.StringVar(&jsonFile, "c", "", "json-file configuration")
	flag.StringVar(&jsonFile, "config", "", "json-file configuration")
	flag.Func("a", "set address", parseStringFlag(&addr))
	flag.Func("p", "poll interval", parseDurationFlag(&pollInterval))
	flag.Func("i", "store interval", parseDurationFlag(&storeInterval))
	flag.Func("f", "store file", parseStringFlag(&storeFile))
	flag.Func("k", "hash secret key", parseStringFlag(&key))
	flag.Func("d", "database url string", parseStringFlag(&dbURL))
	flag.Func("l", "rate limit", parseIntFlag(&rateLimit))
	flag.Func("crypto-key", "set crypto key file (public for agent, private for server)", parseStringFlag(&cryptoKeyPath))
}

// NewAgentConfig - конструктор конфигурации для агента.
func NewAgentConfig() AgentConfig {
	flag.Func("r", "report interval", parseDurationFlag(&reportInterval))
	flag.Parse()
	log.Println(addr)
	parseENV()
	cfg := json.ParseJSON(jsonFile)

	cryptoKey, err := getPublicKey(cfgSwitch(cryptoKeyPath, cfg.CryptoKey))
	if err != nil {
		log.Println("can't get public crypto key")
		log.Fatal(err)
	}

	return AgentConfig{
		Addr:           cfgSwitch(addr, cfg.Address),
		ReportInterval: cfgSwitch(reportInterval, cfg.ReportInterval),
		PoolInterval:   cfgSwitch(pollInterval, cfg.PollInterval),
		HashKey:        cfgSwitch(key, cfg.Key),
		RateLimit:      cfgSwitch(rateLimit, cfg.RateLimit),
		PublicKey:      cryptoKey,
	}
}

// NewServerConfig - конструктор конфигурации для сервера.
func NewServerConfig() ServerConfig {
	flag.Func("r", "is restore from file", parseStringFlag(&restore))
	flag.Parse()
	parseENV()
	cfg := json.ParseJSON(jsonFile)

	cryptoKey, err := getPrivateKey(cfgSwitch(cryptoKeyPath, cfg.CryptoKey))
	if err != nil {
		log.Println("can't get private crypto key")
		log.Fatal(err)
	}

	return ServerConfig{
		Addr:          cfgSwitch(addr, cfg.Address),
		StoreInterval: cfgSwitch(storeInterval, cfg.StoreInterval),
		StoreFile:     cfgSwitch(storeFile, cfg.StoreFile),
		Restore:       cfgSwitch(restore, cfg.Restore) == "true",
		HashKey:       cfgSwitch(key, cfg.Key),
		DBURL:         cfgSwitch(dbURL, cfg.DatabaseDSN),
		PrivateKey:    cryptoKey,
	}
}

func cfgSwitch[T constraints.Ordered](val *T, def T) T {
	if val == nil {
		return def
	}
	return *val
}

func parseDurationFlag(val **time.Duration) func(s string) error {
	return func(s string) error {
		interval, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		*val = &interval
		return nil
	}
}

func parseStringFlag(val **string) func(s string) error {
	return func(s string) error {
		*val = &s
		return nil
	}
}

func parseIntFlag(val **int) func(s string) error {
	return func(i string) (err error) {
		v, err := strconv.Atoi(i)
		if err != nil {
			return err
		}
		*val = &v
		return nil
	}
}
