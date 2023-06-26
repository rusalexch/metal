package poll

import (
	"context"

	"github.com/rusalexch/metal/internal/app"
)

// poller - интерфейс модуля сбора метрик.
type poller interface {
	// ScanToChan - метод сканирование метрик в канал.
	ScanToChan(ctx context.Context, metricCh chan<- app.Metrics)
	// Scan - метод сканирования метрик.
	Scan() []app.Metrics
}
