package services

import (
	"strconv"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/storage"
)

// NewMertricsService конструктор сервиса обработки метрик
func NewMertricsService(storage storage.MetricsStorage) *MertricsService {
	return &MertricsService{
		storage: storage,
	}
}

// Add метод сохранения новой метрики
func (ms *MertricsService) Add(m app.Metric) error {
	switch m.Type {
	case app.Guage:
		return ms.addGuage(m)
	case app.Counter:
		return ms.addCounter(m)
	default:
		return IncorrectTypeErr
	}
}

// Get Метод получения метрики
func (ms *MertricsService) Get(name string, mType app.MetricType) (app.Metric, error) {
	switch mType {
	case app.Guage:
		return ms.getGuage(name)
	case app.Counter:
		return ms.getCounter(name)
	default:
		return app.Metric{}, IncorrectTypeErr
	}
}

// addGuage метода сохранения метрики типа guage
func (ms *MertricsService) addGuage(m app.Metric) error {
	val, err := strconv.ParseFloat(m.Value, 64)
	if err != nil {
		return err
	}
	return ms.storage.AddGuage(m.Name, val)
}

// addCounter метод добавления метрики типа counter
func (ms *MertricsService) addCounter(m app.Metric) error {
	val, err := strconv.ParseInt(m.Value, 10, 64)
	if err != nil {
		return err
	}

	return ms.storage.AddCounter(m.Name, val)
}

// getGuage метод получения метрики типа guage
func (ms *MertricsService) getGuage(name string) (app.Metric, error) {
	var m app.Metric
	val, err := ms.storage.GetGuage(name)
	if err != nil {
		return m, err
	}

	m = app.Metric{
		Type:      app.Guage,
		Value:     strconv.FormatFloat(val, 'E', -1, 64),
		Timestamp: 0,
		Name:      name,
	}

	return m, nil
}

// getCounter метод получения метрики типа counter
func (ms *MertricsService) getCounter(name string) (app.Metric, error) {
	var m app.Metric
	val, err := ms.storage.GetCounter(name)
	if err != nil {
		return m, err
	}

	m = app.Metric{
		Type:      app.Counter,
		Value:     strconv.FormatInt(val, 10),
		Timestamp: 0,
		Name:      name,
	}

	return m, nil
}