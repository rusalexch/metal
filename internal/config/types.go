package config

import "time"

type AgentConfig struct {
	Addr           string        `env:"ADDRESS"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
	PoolInterval   time.Duration `env:"POLL_INTERVAL"`
	HashKey        string        `env:"KEY"`
}

type ServerConfig struct {
	Addr          string        `env:"ADDRESS"`
	StoreInterval time.Duration `env:"STORE_INTERVAL"`
	StoreFile     string        `env:"STORE_FILE"`
	Restore       bool          `env:"RESTORE"`
	HashKey       string        `env:"KEY"`
}
