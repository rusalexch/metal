package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

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
