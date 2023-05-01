package poll

import (
	"context"
	"time"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/poll/internal/gopsutil"
	"github.com/rusalexch/metal/internal/poll/internal/runtime"
)

// poll структура настроек модуля metrics
type poll struct {
	rt        poller
	rtTrigger chan interface{}
	gu        poller
	guTrigger chan interface{}
}

// New создание модуля сбора метрик
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

func (p *poll) ScanChan(ctx context.Context, ticker *time.Ticker, metricCh chan<- app.Metrics) {
	go func() {
		rtCtx, rtCancel := context.WithCancel(ctx)
		guCtx, guCancel := context.WithCancel(ctx)
		for {
			select {
			case <-ctx.Done():
				p.close()
				rtCancel()
				guCancel()
				return
			case <-ticker.C:
				p.rt.ScanToChan(rtCtx, metricCh)
				p.gu.ScanToChan(guCtx, metricCh)
			}
		}
	}()
}

func (p *poll) close() {
	if p.rtTrigger != nil {
		close(p.rtTrigger)
	}
	if p.guTrigger != nil {
		close(p.rtTrigger)
	}
}

func (p *poll) Scan() []app.Metrics {
	rtM := p.rt.Scan()
	guM := p.gu.Scan()
	res := make([]app.Metrics, 0, len(rtM)+len(guM))
	res = append(res, rtM...)
	res = append(res, guM...)

	return res
}
