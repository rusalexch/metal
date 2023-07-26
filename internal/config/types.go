package config

import (
	"crypto/rsa"
	"time"
)

// AgentConfig - структура конфигурации агента.
type AgentConfig struct {
	// адрес сервера.
	Addr string `env:"ADDRESS"`
	// ключ хэш-функции.
	HashKey string `env:"KEY"`
	// интервал сбора метрик.
	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
	// интервал отправки метрик.
	PoolInterval time.Duration `env:"POLL_INTERVAL"`
	// количество одновременно исходящих запросов от агента.
	RateLimit int `env:"RATE_LIMIT"`
	// PublicKey публичный ключ
	PublicKey *rsa.PublicKey
}

// ServerConfig - структура конфигурации сервера.
type ServerConfig struct {
	// адрес сервера.
	Addr string `env:"ADDRESS"`
	// путь к файлу файлового хранилища.
	StoreFile string `env:"STORE_FILE"`
	// ключ хэш-функции.
	HashKey string `env:"KEY"`
	// url подключения к базе данных.
	DBURL string `env:"DATABASE_DSN"`
	// интервал сохранения данных в файловое хранилище.
	StoreInterval time.Duration `env:"STORE_INTERVAL"`
	// флаг восстановления данных из файла файлового хранилища.
	Restore bool `env:"RESTORE"`
	// PrivateKey приватный ключ
	PrivateKey *rsa.PrivateKey
}

// type defaultValues struct {
// 		// адрес сервера по умолчанию.
// 		defaultAddr string
// 		// интервал сбора метрик по умолчанию.
// 		defaultReportInterval time.Duration
// 		// интервал отправки метрик по умолчанию.
// 		defaultPoolInterval time.Duration
// 		// статус восстановления метрик из файлового хранилища по умолчанию.
// 		defaultRestore = "true"
// 		// интервал сохранения метрик в файловое хранилище по умолчанию.
// 		defaultStoreInterval = time.Second * 300
// 		// путь к файлу файлового хранилища по умолчанию.
// 		defaultStoreFile = "/tmp/devops-metrics-db.json"
// 		// ключ хэш-функции по умолчанию.
// 		defaultKey = ""
// 		// количество одновременно исходящих запросов по умолчанию.
// 		defaultRateLimit = 1
// }


