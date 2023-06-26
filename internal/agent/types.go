package agent

import (
	"time"
)

// Agent структура настроек приложения
type Agent struct {
	pollInterval   time.Duration // pollInterval - частота сбора метрик
	reportInterval time.Duration // reportInterval - частота отправки метрик на сервер
	poll           Scaner        // poll -  пакет сбора метрик
	cache          Cache         // cache - пакет сохранения метрик в памяти
	transport      Transport     // client - клиент для отправки метрик на сервер
	hash           hasher        // hash - хеш-пакет
}

// Config конфигурация приложения
type Config struct {
	PollInterval   time.Duration // PollInterval - частота сбора метрик
	ReportInterval time.Duration // ReportInterval - частота отправки метрик на сервер
	Poll           Scaner        // Poll - пакет сбора метрик
	Cache          Cache         // Cache - пакет сохранения метрик в памяти
	Transport      Transport     // Transport - клиент для отправки метрик на сервер
	Hash           hasher        // Hash - хеш-пакет
}
