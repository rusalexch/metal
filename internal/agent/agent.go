package agent

import (
	"context"
	"time"

	"github.com/rusalexch/metal/internal/app"
)

// New инициализация приложения.
func New(conf Config) *Agent {

	return &Agent{
		pollInterval:   conf.PollInterval,
		reportInterval: conf.ReportInterval,
		poll:           conf.Poll,
		cache:          conf.Cache,
		transport:      conf.Transport,
		hash:           conf.Hash,
	}
}

// Start метод запуска клиента сбора и отправки метрик на сервер.
func (a *Agent) Start(ctx context.Context) {
	pollTicker := time.NewTicker(a.pollInterval)
	defer pollTicker.Stop()
	reportTicker := time.NewTicker(a.reportInterval)
	defer reportTicker.Stop()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	pollChan := make(chan app.Metrics)
	defer close(pollChan)
	reqChan := make(chan []app.Metrics)
	defer close(reqChan)

	a.poll.ScanChan(ctx, pollTicker, pollChan)
	a.cache.Start(ctx, pollChan, reqChan, *reportTicker)
	a.transport.Start(ctx, reqChan)
}
