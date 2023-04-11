package services

import "github.com/rusalexch/metal/internal/storage"

// New конструктор сервисов
func New(storage storage.MetricsStorage) *Services {
	ms := NewMertricsService(storage)
	h := NewHealthCheck(storage)

	return &Services{
		MetricsService: ms,
		HealthCheck:    h,
	}
}
