package agent

import (
	"time"
)

// Agent структура настроек приложения.
type Agent struct {
	// poll -  пакет сбора метрик.
	poll Scaner
	// cache - пакет сохранения метрик в памяти.
	cache Cache
	// client - клиент для отправки метрик на сервер.
	transport Transport
	// hash - хеш-пакет.
	hash hasher
	// pollInterval - частота сбора метрик.
	pollInterval time.Duration
	// reportInterval - частота отправки метрик на сервер.
	reportInterval time.Duration
}

// Config конфигурация приложения.
type Config struct {
	// Poll - пакет сбора метрик.
	Poll Scaner
	// Cache - пакет сохранения метрик в памяти.
	Cache Cache
	// Transport - клиент для отправки метрик на сервер.
	Transport Transport
	// Hash - хеш-пакет.
	Hash hasher
	// PollInterval - частота сбора метрик.
	PollInterval time.Duration
	// ReportInterval - частота отправки метрик на сервер.
	ReportInterval time.Duration
}
