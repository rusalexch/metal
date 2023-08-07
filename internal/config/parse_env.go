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
		reportInterval = parseDurationEnv(reportIntervalEnv)
	}
	if poolIntervalEnv, isSet := os.LookupEnv("POLL_INTERVAL"); isSet {
		pollInterval = parseDurationEnv(poolIntervalEnv)
	}
	if storeIntervalEnv, isSet := os.LookupEnv("STORE_INTERVAL"); isSet {
		storeInterval = parseDurationEnv(storeIntervalEnv)
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
		rateLimit = parseIntEnv(rateLimitEnv)
	}
	if cryptoKeyPathEnv, isSet := os.LookupEnv("CRYPTO_KEY"); isSet {
		cryptoKeyPath = &cryptoKeyPathEnv
	}
	if jsonFileEnv, isSet := os.LookupEnv("CONFIG"); isSet {
		jsonFile = jsonFileEnv
	}
	if trustedSubnetEnv, isSet := os.LookupEnv("TRUSTED_SUBNET"); isSet {
		trustedSubnet = &trustedSubnetEnv
	}
	if grpcAddressEnv, isSet := os.LookupEnv("GRPC_ADDRESS"); isSet {
		grpcAddress = &grpcAddressEnv
	}
}

func parseDurationEnv(s string) *time.Duration {
	t, err := time.ParseDuration(s)
	if err != nil {
		log.Fatal(err)
	}

	return &t
}

func parseIntEnv(s string) *int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return &num
}
