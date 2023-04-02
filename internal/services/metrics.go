package services

import (
	"log"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/storage"
)

// NewMertricsService конструктор сервиса обработки метрик
func NewMertricsService(storage storage.MetricsStorage) *MertricsService {
	return &MertricsService{
		storage:     storage,
		subscribers: []func(){},
	}
}

// Add метод сохранения новой метрики
func (ms *MertricsService) Add(m app.Metrics) (err error) {
	switch m.Type {
	case app.Gauge:
		err = ms.addGuage(m)
	case app.Counter:
		err = ms.addCounter(m)
	default:
		err = ErrIncorrectType
	}
	if err == nil {
		ms.eventAdd()
	}
	return
}

// Get Метод получения метрики
func (ms *MertricsService) Get(name string, mType app.MetricType) (app.Metrics, error) {
	switch mType {
	case app.Gauge:
		return ms.getGuage(name)
	case app.Counter:
		return ms.getCounter(name)
	default:
		return app.Metrics{}, ErrIncorrectType
	}
}

func (ms *MertricsService) List() []app.Metrics {
	counters := ms.storage.ListCounter()
	gauges := ms.storage.ListGauge()

	res := make([]app.Metrics, 0, len(counters)+len(gauges))
	for _, val := range counters {
		res = append(res, app.Metrics{
			Type:  app.Counter,
			Delta: &val.Value,
			ID:    val.Name,
		})
	}
	for _, val := range gauges {
		res = append(res, app.Metrics{
			Type:  app.Gauge,
			Value: &val.Value,
			ID:    val.Name,
		})
	}

	return res
}

func (ms *MertricsService) Subscribe(f func()) {
	if ms.subscribers != nil {
		ms.subscribers = append(ms.subscribers, f)
	}
	ms.subscribers = []func(){f}

}

// addGuage метода сохранения метрики типа guage
func (ms *MertricsService) addGuage(m app.Metrics) error {
	return ms.storage.AddGauge(m.ID, *m.Value)
}

// addCounter метод добавления метрики типа counter
func (ms *MertricsService) addCounter(m app.Metrics) error {
	return ms.storage.AddCounter(m.ID, *m.Delta)
}

// getGuage метод получения метрики типа guage
func (ms *MertricsService) getGuage(name string) (app.Metrics, error) {
	var m app.Metrics
	val, err := ms.storage.GetGauge(name)
	if err != nil {
		return m, err
	}

	m = app.Metrics{
		Type:  app.Gauge,
		Value: &val,
		ID:    name,
	}

	return m, nil
}

// getCounter метод получения метрики типа counter
func (ms *MertricsService) getCounter(name string) (app.Metrics, error) {
	val, err := ms.storage.GetCounter(name)
	if err != nil {
		return app.Metrics{}, err
	}

	m := app.Metrics{
		Type:  app.Counter,
		Delta: &val,
		ID:    name,
	}

	return m, nil
}

func (ms *MertricsService) eventAdd() {
	log.Println(ms.subscribers)
	if ms.subscribers != nil && len(ms.subscribers) > 0 {
		for _, f := range ms.subscribers {
			f()
		}
	}
}
