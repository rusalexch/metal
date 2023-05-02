package poll

import (
	"context"

	"github.com/rusalexch/metal/internal/app"
)

type poller interface {
	ScanToChan(ctx context.Context, metricCh chan<- app.Metrics)
	Scan() []app.Metrics
}
