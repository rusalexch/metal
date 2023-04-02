package config

import "time"

type AgentConfig struct {
	Addr           string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
	PoolInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
}

type ServerConfig struct {
	Addr          string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"300s"`
	StoreFile     string        `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	Restore       bool          `env:"RESTORE" envDefault:"true"`
}
