package config

import "time"

// AgentConfig - структура конфигурации агента.
type AgentConfig struct {
	// адрес сервера.
	Addr string `env:"ADDRESS"`
	// интервал сбора метрик.
	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
	// интервал отправки метрик.
	PoolInterval time.Duration `env:"POLL_INTERVAL"`
	// ключ хэш-функции.
	HashKey string `env:"KEY"`
	// количество одновременно исходящих запросов от агента.
	RateLimit int `env:"RATE_LIMIT"`
}

// ServerConfig - структура конфигурации сервера.
type ServerConfig struct {
	// адрес сервера.
	Addr string `env:"ADDRESS"`
	// интервал сохранения данных в файловое хранилище.
	StoreInterval time.Duration `env:"STORE_INTERVAL"`
	// путь к файлу файлового хранилища.
	StoreFile string `env:"STORE_FILE"`
	// флаг восстановления данных из файла файлового хранилища.
	Restore bool `env:"RESTORE"`
	// ключ хэш-функции.
	HashKey string `env:"KEY"`
	// url подключения к базе данных.
	DBURL string `env:"DATABASE_DSN"`
}
