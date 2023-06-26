package poll

import (
	"context"
	"time"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/poll/internal/gopsutil"
	"github.com/rusalexch/metal/internal/poll/internal/runtime"
)

// poll - структура настроек модуля metrics
type poll struct {
	// rt - модуль сбора метрик runtime
	rt poller
	// rtTrigger - триггер сбора метрик runtime
	rtTrigger chan interface{}
	// gu - модуль сбора метрик gopsutil
	gu poller
	// guTrigger - триггер сбора метрик gopsutil
	guTrigger chan interface{}
}

// New - конструктор модуля сбора метрик
func New() *poll {
	rtTrigger := make(chan interface{})
	guTrigger := make(chan interface{})
	return &poll{
		rt:        runtime.New(rtTrigger),
		rtTrigger: rtTrigger,
		gu:        gopsutil.New(guTrigger),
		guTrigger: guTrigger,
	}
}

// ScanChan - метод сканирования метрик в каналы
func (p *poll) ScanChan(ctx context.Context, ticker *time.Ticker, metricCh chan<- app.Metrics) {
	go func() {
		rtCtx, rtCancel := context.WithCancel(context.Background())
		guCtx, guCancel := context.WithCancel(context.Background())
		p.rt.ScanToChan(rtCtx, metricCh)
		p.gu.ScanToChan(guCtx, metricCh)
		for {
			select {
			case <-ctx.Done():
				rtCancel()
				guCancel()
				p.close()
				return
			case <-ticker.C:
				p.rtTrigger <- struct{}{}
				p.guTrigger <- struct{}{}
			}
		}
	}()
}

// close - метод закрытия модуля сбора метрик
func (p *poll) close() {
	if p.rtTrigger != nil {
		close(p.rtTrigger)
	}
	if p.guTrigger != nil {
		close(p.rtTrigger)
	}
}

// Scan - метод сканирования метрик
func (p *poll) Scan() []app.Metrics {
	rtM := p.rt.Scan()
	guM := p.gu.Scan()
	res := make([]app.Metrics, 0, len(rtM)+len(guM))
	res = append(res, rtM...)
	res = append(res, guM...)

	return res
}
