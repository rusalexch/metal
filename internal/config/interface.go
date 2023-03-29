package config

type AgentConfig struct {
	Addr           string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	ReportInterval int    `env:"REPORT_INTERVAL" envDefault:"10"`
	PoolInterval   int    `env:"POLL_INTERVAL" envDefault:"2"`
}

type ServerConfig struct {
	Addr string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
}
