package app

import "time"

// MetricType тип для типа метрики
type MetricType = string

// Metric структура метрики
type Metric struct {
	// Type наименование типа метрики
	Type MetricType
	// Value значение метрики в строковом формате
	Value string
	// Timestamp дата и время сбора метрики
	Timestamp int64
	// Name наименование метрики
	Name string
}

// Server структура настроек сервера
type Server struct {
	// url адрес сервера сбора метрик
	url string
	// port порт сервера сбора метрик
	port int
}

// App структура настроек приложения
type App struct {
	// pollInterval частота сбора метрик
	pollInterval time.Duration
	// reportInterval частота отправки метрик на сервер
	reportInterval time.Duration
	// // server конфигурация сервера сбора метрик
	// server Server
	// metrics пакет сбора метрик
	metrics Metrics
	// cache пакет сохранения метрик в памяти
	cache Cache
	// client клиент для отправки метрик на сервер
	transport Transport
}

// Config конфигурация приложения
type Config struct {
	// pollInterval частота сбора метрик
	PollInterval time.Duration
	// reportInterval частота отправки метрик на сервер
	ReportInterval time.Duration
	// server адрес сервера сбора метрик
	ServerUrl string
	// server порт сервера сбора метрик
	ServerPort int
	// metrics пакет сбора метрик
	Metrics Metrics
	// cache пакет сохранения метрик в памяти
	Cache Cache
	// client клиент для отправки метрик на сервер
	Transport Transport
}
