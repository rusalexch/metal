package agent

import (
	"context"
	"time"

	"github.com/rusalexch/metal/internal/app"
)

// Scaner - интерфейс для сканера метрик.
type Scaner interface {
	// Scan - метод для сканирования метрик.
	Scan() []app.Metrics
	// ScanChan - метод сканирования метрик в канал.
	ScanChan(ctx context.Context, ticker *time.Ticker, metricCh chan<- app.Metrics)
}

// Transport - интерфейс клиента для отправки метрик на сервер.
type Transport interface {
	// Start - метод запуска клиента отправки метрик.
	Start(ctx context.Context, ch <-chan []app.Metrics)
}

// Cache - интерфейс кэша для хранения метрик.
type Cache interface {
	// Start - Метода запуска сервиса кэширования.
	Start(ctx context.Context, chIn <-chan app.Metrics, chOut chan<- []app.Metrics, t time.Ticker)
}

// hasher = интерфейс дял хэш-функции.
type hasher interface {
	// AddHash - метод добавления хэша к метрики.
	AddHash(m *app.Metrics)
	// Check - метод проверки хэша метрики.
	Check(m app.Metrics) bool
}
