package config

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	addr           *string
	reportInterval time.Duration
	pollInterval   time.Duration
	storeInterval  time.Duration
	storeFile      *string
	restore        *string
	key            *string
	dbURL          *string
	rateLimit      int
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
}

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
	return AgentConfig{
		Addr:           *addr,
		ReportInterval: reportInterval,
		PoolInterval:   pollInterval,
		HashKey:        *key,
		RateLimit:      rateLimit,
	}
}

func NewServerConfig() ServerConfig {
	restore = flag.String("r", defaultRestore, "is restore from file")
	flag.Parse()
	parseENV()
	return ServerConfig{
		Addr:          *addr,
		StoreInterval: storeInterval,
		StoreFile:     *storeFile,
		Restore:       *restore == "true",
		HashKey:       *key,
		DBURL:         *dbURL,
	}
}

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
}
