package handlers

import (
	"context"

	"github.com/rusalexch/metal/internal/app"
)

// storager - интерфейс хранилища
type storager interface {
	// Add - добавить метрику в хранилище
	Add(ctx context.Context, m app.Metrics) error
	// AddList - добавить список метрик
	AddList(ctx context.Context, m []app.Metrics) error
	// Get - получение метрики по имени и типу
	Get(ctx context.Context, name string, mType app.MetricType) (app.Metrics, error)
	// List - получение всех метрик списком
	List(ctx context.Context) ([]app.Metrics, error)
	// Ping - проверка связи с хранилищем
	Ping(ctx context.Context) error
	// Close - метод закрытия сессии с хранилищем
	Close()
}

// hasher - интерфейс хэш-функции
type hasher interface {
	// AddHash - метод добавления хэша к метрики
	AddHash(m *app.Metrics)
	// Check - метод проверки хеша метрики
	Check(m app.Metrics) bool
}
