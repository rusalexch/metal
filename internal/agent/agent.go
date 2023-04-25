package agent

import (
	"context"
	"time"

	"github.com/rusalexch/metal/internal/app"
)

// New инициализация приложения
// pollInterval - частота опроса метрик
// reportInterval - частота отправки метрик на сервер
// url - адрес сервера, по умолчанию "http://127.0.0.1"
// port - порт сервера, по умолчанию 8080
func New(conf Config) *Agent {

	return &Agent{
		pollInterval:   conf.PollInterval,
		reportInterval: conf.ReportInterval,
		metrics:        conf.Metrics,
		cache:          conf.Cache,
		transport:      conf.Transport,
		hash:           conf.Hash,
	}
}

// Start метод запуска клиента сбора и отправки метрик на сервер
func (a *Agent) Start() {
	pollTicker := time.NewTicker(a.pollInterval)
	defer pollTicker.Stop()
	reportTicker := time.NewTicker(a.reportInterval)
	defer reportTicker.Stop()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pollChan := make(chan app.Metrics)
	defer close(pollChan)
	reqChan := make(chan []app.Metrics)
	defer close(reqChan)

	a.metrics.ScanChan(ctx, pollTicker, pollChan)
	a.cache.Start(ctx, pollChan, reqChan, *reportTicker)
	a.transport.Start(ctx, reqChan)
}
