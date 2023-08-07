package grpcserver

import (
	"context"

	"github.com/rusalexch/metal/internal/app"
)

// metricsStorage - интерфейс хранилища метрик.
type storage interface {
	// Add - метод добавления/обновления метрики.
	Add(ctx context.Context, m app.Metrics) error
}
