package config

import (
	"flag"
	"log"
	"os"
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
)

func init() {
	addr = flag.String("a", defaultAddr, "set address")
	pollInterval = defaultPoolInterval
	flag.Func("p", "poool interval", func(s string) (err error) {
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
	checkENV()
	return AgentConfig{
		Addr:           *addr,
		ReportInterval: reportInterval,
		PoolInterval:   pollInterval,
		HashKey:        *key,
	}
}

func NewServerConfig() ServerConfig {
	restore = flag.String("r", defaultRestore, "is restore from file")
	flag.Parse()
	checkENV()
	return ServerConfig{
		Addr:          *addr,
		StoreInterval: storeInterval,
		StoreFile:     *storeFile,
		Restore:       *restore == "true",
		HashKey:       *key,
	}
}

func checkENV() {
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
}
