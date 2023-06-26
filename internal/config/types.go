package config

import "time"

// AgentConfig - структура конфигурации агента
type AgentConfig struct {
	Addr           string        `env:"ADDRESS"`         // адрес сервера
	ReportInterval time.Duration `env:"REPORT_INTERVAL"` // интервал сбора метрик
	PoolInterval   time.Duration `env:"POLL_INTERVAL"`   // интервал отправки метрик
	HashKey        string        `env:"KEY"`             // ключ хэш-функции
	RateLimit      int           `env:"RATE_LIMIT"`      // количество одновременно исходящих запросов от агента
}

// ServerConfig - структура конфигурации сервера
type ServerConfig struct {
	Addr          string        `env:"ADDRESS"`        // адрес сервера
	StoreInterval time.Duration `env:"STORE_INTERVAL"` // интервал сохранения данных в файловое хранилище
	StoreFile     string        `env:"STORE_FILE"`     // путь к файлу файлового хранилища
	Restore       bool          `env:"RESTORE"`        // флаг восстановления данных из файла файлового хранилища
	HashKey       string        `env:"KEY"`            // ключ хэш-функции
	DBURL         string        `env:"DATABASE_DSN"`   // url подключения к базе данных
}
