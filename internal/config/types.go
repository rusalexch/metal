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
	// CryptoKey публичный ключ
	CryptoKey *rsa.PublicKey
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
	// CryptoKey приватный ключ
	CryptoKey *rsa.PrivateKey
}
