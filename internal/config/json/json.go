package json

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"
)

type DefaultConfig struct {
	Address        string
	Restore        string
	StoreInterval  time.Duration
	StoreFile      string
	DatabaseDSN    string
	CryptoKey      string
	ReportInterval time.Duration
	PollInterval   time.Duration
	Key            string
	RateLimit      int
	TrustedSubnet  string
	GRPCAddress    string
}

type JSONConfig struct {
	Address        *string   `json:"address,omitempty"`
	Restore        *string   `json:"restore,omitempty"`
	StoreInterval  *interval `json:"store_interval,omitempty"`
	StoreFile      *string   `json:"store_file,omitempty"`
	DatabaseDSN    *string   `json:"database_dsn,omitempty"`
	CryptoKey      *string   `json:"crypto_key,omitempty"`
	ReportInterval *interval `json:"report_interval,omitempty"`
	PollInterval   *interval `json:"poll_interval,omitempty"`
	TrustedSubnet  *string   `json:"trusted_subnet,omitempty"`
	GRPCAddress    *string   `json:"grpc_address,omitempty"`
}

func ParseJSON(jsonFile string) *DefaultConfig {
	if jsonFile == "" {
		return defaultValues()
	}
	file, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var jsonCfg JSONConfig
	err = json.Unmarshal(file, &jsonCfg)
	if err != nil {
		log.Fatal(err)
	}

	return fillDefaultValues(jsonCfg)
}

func defaultValues() *DefaultConfig {
	return &DefaultConfig{
		Address:        defaultAddr,
		Restore:        defaultRestore,
		StoreInterval:  defaultStoreInterval,
		StoreFile:      defaultStoreFile,
		DatabaseDSN:    "",
		CryptoKey:      "",
		ReportInterval: defaultReportInterval,
		PollInterval:   defaultPoolInterval,
		Key:            defaultKey,
		RateLimit:      defaultRateLimit,
		TrustedSubnet:  defaultTrustedSubnet,
		GRPCAddress:    defaultGRPCAddress,
	}
}

func fillDefaultValues(json JSONConfig) *DefaultConfig {
	defValues := defaultValues()
	if json.Address != nil {
		defValues.Address = *json.Address
	}
	if json.CryptoKey != nil {
		defValues.CryptoKey = *json.CryptoKey
	}
	if json.DatabaseDSN != nil {
		defValues.DatabaseDSN = *json.DatabaseDSN
	}
	if json.PollInterval != nil {
		defValues.PollInterval = json.PollInterval.Duration
	}
	if json.ReportInterval != nil {
		defValues.ReportInterval = json.ReportInterval.Duration
	}
	if json.Restore != nil {
		defValues.Restore = *json.Restore
	}
	if json.StoreFile != nil {
		defValues.StoreFile = *json.StoreFile
	}
	if json.StoreInterval != nil {
		defValues.StoreInterval = json.StoreInterval.Duration
	}
	if json.TrustedSubnet != nil {
		defValues.TrustedSubnet = *json.TrustedSubnet
	}
	if json.GRPCAddress != nil {
		defValues.GRPCAddress = *json.GRPCAddress
	}
	return defValues
}

type interval struct {
	time.Duration
}

func (i interval) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *interval) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		i.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		i.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}
