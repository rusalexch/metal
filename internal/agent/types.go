package agent

import (
	"time"
)

// Agent структура настроек приложения
type Agent struct {
	// pollInterval частота сбора метрик
	pollInterval time.Duration
	// reportInterval частота отправки метрик на сервер
	reportInterval time.Duration
	// metrics пакет сбора метрик
	metrics Metrics
	// cache пакет сохранения метрик в памяти
	cache Cache
	// client клиент для отправки метрик на сервер
	transport Transport
	// hash хеш-пакет
	hash hasher
}

// Config конфигурация приложения
type Config struct {
	// pollInterval частота сбора метрик
	PollInterval time.Duration
	// reportInterval частота отправки метрик на сервер
	ReportInterval time.Duration
	// metrics пакет сбора метрик
	Metrics Metrics
	// cache пакет сохранения метрик в памяти
	Cache Cache
	// client клиент для отправки метрик на сервер
	Transport Transport
	// Hash  хеш-пакет
	Hash hasher
}
