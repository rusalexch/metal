package gopsutil

import (
	"context"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"

	"github.com/rusalexch/metal/internal/app"
)

// gopsutil - структура модуля метрик gopsutil
type gopsutil struct {
	trigger <-chan interface{}
}

// New - конструктор модула метрик gopsutil
func New(trigger <-chan interface{}) *gopsutil {
	return &gopsutil{
		trigger: trigger,
	}
}

// ScanToChan - метод сканирования метрик gopsutil в канал
func (g *gopsutil) ScanToChan(ctx context.Context, metricCh chan<- app.Metrics) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-g.trigger:
				g.scanToChan(metricCh)
			}
		}
	}()
}

// scanToChan - сканирование метрик gopsutil в канал
func (g *gopsutil) scanToChan(metricCh chan<- app.Metrics) {
	for _, v := range g.Scan() {
		metricCh <- v
	}
}

// Scan - метод сканирования метрик gopsutil
func (g *gopsutil) Scan() []app.Metrics {
	m, _ := mem.VirtualMemory()
	c, _ := cpu.Percent(0, false)
	var util float64
	for _, u := range c {
		util += u
	}
	util = util / float64(len(c))

	return []app.Metrics{
		app.AsGauge(float64(m.Total), "TotalMemory"),
		app.AsGauge(float64(m.Free), "FreeMemory"),
		app.AsGauge(float64(util), "CPUutilization1"),
	}
}
