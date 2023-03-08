package metric

import (
	"math/rand"
	"runtime"
	"strconv"
	"time"

	"github.com/rusalexch/metal/internal/app"
)

// New создание модуля сбора метрик
func New() *Metrics {
	return &Metrics{}
}

// Scan сканирование метрики
func (m *Metrics) Scan() ([]app.Metric, error) {
	rm := runtime.MemStats{}
	runtime.ReadMemStats(&rm)
	res := make([]app.Metric, 0, 29)
	m.cnt += 1

	res = append(res, guage(float64(rm.Alloc), "Alloc"))
	res = append(res, guage(float64(rm.BuckHashSys), "BuckHashSys"))
	res = append(res, guage(float64(rm.Frees), "Frees"))
	res = append(res, guage(rm.GCCPUFraction, "GCCPUFraction"))
	res = append(res, guage(float64(rm.GCSys), "GCSys"))
	res = append(res, guage(float64(rm.HeapAlloc), "HeapAlloc"))
	res = append(res, guage(float64(rm.HeapIdle), "HeapIdle"))
	res = append(res, guage(float64(rm.HeapInuse), "HeapInuse"))
	res = append(res, guage(float64(rm.HeapObjects), "HeapObjects"))
	res = append(res, guage(float64(rm.HeapReleased), "HeapReleased"))
	res = append(res, guage(float64(rm.HeapSys), "HeapSys"))
	res = append(res, guage(float64(rm.LastGC), "LastGC"))
	res = append(res, guage(float64(rm.Lookups), "Lookups"))
	res = append(res, guage(float64(rm.MCacheInuse), "MCacheInuse"))
	res = append(res, guage(float64(rm.MCacheSys), "MCacheSys"))
	res = append(res, guage(float64(rm.MSpanInuse), "MSpanInuse"))
	res = append(res, guage(float64(rm.MSpanSys), "MSpanSys"))
	res = append(res, guage(float64(rm.Mallocs), "Mallocs"))
	res = append(res, guage(float64(rm.NextGC), "NextGC"))
	res = append(res, guage(float64(rm.NumForcedGC), "NumForcedGC"))
	res = append(res, guage(float64(rm.NumGC), "NumGC"))
	res = append(res, guage(float64(rm.OtherSys), "OtherSys"))
	res = append(res, guage(float64(rm.PauseTotalNs), "PauseTotalNs"))
	res = append(res, guage(float64(rm.StackInuse), "StackInuse"))
	res = append(res, guage(float64(rm.StackInuse), "StackInuse"))
	res = append(res, guage(float64(rm.StackSys), "StackSys"))
	res = append(res, guage(float64(rm.Sys), "Sys"))
	res = append(res, guage(float64(rm.TotalAlloc), "TotalAlloc"))
	res = append(res, counter(m.cnt, "PollCount"))
	res = append(res, guage(randomValue(), "RandomValue"))

	return res, nil
}

// guage преобразование метрики типа goage
func guage(v float64, name string) app.Metric {
	return app.Metric{
		Type:      app.Guage,
		Value:     strconv.FormatFloat(v, 'E', -1, 64),
		Timestamp: time.Now().Unix(),
		Name:      name,
	}
}

// counter преобразование метрики типа counter
func counter(v int64, name string) app.Metric {
	return app.Metric{
		Type:  app.Counter,
		Value: strconv.FormatInt(v, 16),
		Name:  name,
	}
}

// randomValue получение случайного значения float64 в диапазоне от -100 до 100
func randomValue() float64 {
	var min float64 = -100
	var max float64 = 100
	r := min + rand.Float64()*(max-min)

	return r
}
