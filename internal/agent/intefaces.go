package agent

import (
	"context"
	"time"

	"github.com/rusalexch/metal/internal/app"
)

// Scaner интерфейс для сканера метрик
type Scaner interface {
	// Scan метод для сканирования метрик
	Scan() []app.Metrics
	ScanChan(ctx context.Context, ticker *time.Ticker, metricCh chan<- app.Metrics)
}

// Transport интерфейс клиента для отправки метрик на сервер
type Transport interface {
	Start(ctx context.Context, ch <-chan []app.Metrics)
}

// Cache интерфейс кэша для хранения метрик
type Cache interface {
	Start(ctx context.Context, chIn <-chan app.Metrics, chOut chan<- []app.Metrics, t time.Ticker)
}

type hasher interface {
	AddHash(m *app.Metrics)
	Check(m app.Metrics) bool
}
