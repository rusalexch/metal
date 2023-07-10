package runtime

import (
	"context"
	"math/rand"
	rt "runtime"

	"github.com/rusalexch/metal/internal/app"
)

// runtime - структура модуля сбора метрик runtime
type runtime struct {
	trigger <-chan interface{}
	cnt     int64
}

// New - конструктор модуля метрик runtime
func New(trigger <-chan interface{}) *runtime {
	return &runtime{
		trigger: trigger,
	}
}

// ScanToChan - метод сканирование метрик runtime в канал
func (r *runtime) ScanToChan(ctx context.Context, metricCh chan<- app.Metrics) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-r.trigger:
				r.scanToChan(metricCh)
			}
		}
	}()
}

// ScanToChan - сканирование метрик runtime в канал
func (r *runtime) scanToChan(metricCh chan<- app.Metrics) {
	for _, v := range r.Scan() {
		metricCh <- v
	}
}

// Scan - метод сканирование метрики runtime
func (r *runtime) Scan() []app.Metrics {
	rm := rt.MemStats{}
	rt.ReadMemStats(&rm)
	res := make([]app.Metrics, 0, 29)
	r.cnt += 1

	res = append(res, app.AsGauge(float64(rm.Alloc), "Alloc"))
	res = append(res, app.AsGauge(float64(rm.BuckHashSys), "BuckHashSys"))
	res = append(res, app.AsGauge(float64(rm.Frees), "Frees"))
	res = append(res, app.AsGauge(rm.GCCPUFraction, "GCCPUFraction"))
	res = append(res, app.AsGauge(float64(rm.GCSys), "GCSys"))
	res = append(res, app.AsGauge(float64(rm.HeapAlloc), "HeapAlloc"))
	res = append(res, app.AsGauge(float64(rm.HeapIdle), "HeapIdle"))
	res = append(res, app.AsGauge(float64(rm.HeapInuse), "HeapInuse"))
	res = append(res, app.AsGauge(float64(rm.HeapObjects), "HeapObjects"))
	res = append(res, app.AsGauge(float64(rm.HeapReleased), "HeapReleased"))
	res = append(res, app.AsGauge(float64(rm.HeapSys), "HeapSys"))
	res = append(res, app.AsGauge(float64(rm.LastGC), "LastGC"))
	res = append(res, app.AsGauge(float64(rm.Lookups), "Lookups"))
	res = append(res, app.AsGauge(float64(rm.MCacheInuse), "MCacheInuse"))
	res = append(res, app.AsGauge(float64(rm.MCacheSys), "MCacheSys"))
	res = append(res, app.AsGauge(float64(rm.MSpanInuse), "MSpanInuse"))
	res = append(res, app.AsGauge(float64(rm.MSpanSys), "MSpanSys"))
	res = append(res, app.AsGauge(float64(rm.Mallocs), "Mallocs"))
	res = append(res, app.AsGauge(float64(rm.NextGC), "NextGC"))
	res = append(res, app.AsGauge(float64(rm.NumForcedGC), "NumForcedGC"))
	res = append(res, app.AsGauge(float64(rm.NumGC), "NumGC"))
	res = append(res, app.AsGauge(float64(rm.OtherSys), "OtherSys"))
	res = append(res, app.AsGauge(float64(rm.PauseTotalNs), "PauseTotalNs"))
	res = append(res, app.AsGauge(float64(rm.StackInuse), "StackInuse"))
	res = append(res, app.AsGauge(float64(rm.StackInuse), "StackInuse"))
	res = append(res, app.AsGauge(float64(rm.StackSys), "StackSys"))
	res = append(res, app.AsGauge(float64(rm.Sys), "Sys"))
	res = append(res, app.AsGauge(float64(rm.TotalAlloc), "TotalAlloc"))
	res = append(res, app.AsCounter(r.cnt, "PollCount"))
	res = append(res, app.AsGauge(randomValue(), "RandomValue"))

	return res
}

// randomValue - получение случайного значения float64 в диапазоне от -100 до 100
func randomValue() float64 {
	var min float64 = -100
	var max float64 = 100
	r := min + rand.Float64()*(max-min)

	return r
}
