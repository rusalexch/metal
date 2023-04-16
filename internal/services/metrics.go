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
	return ms.storage.Add(m)
}

// Get Метод получения метрики
func (ms *MertricsService) Get(name string, mType app.MetricType) (app.Metrics, error) {
	m, err := ms.storage.Get(name, mType)
	if err != nil {
		return app.Metrics{}, err
	}
	return *m, nil
}

func (ms *MertricsService) List() []app.Metrics {
	l, err := ms.storage.List()
	if err != nil {
		log.Println(err)
		return []app.Metrics{}
	}

	return l
}
