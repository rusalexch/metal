package storage

import (
	"context"

	"github.com/rusalexch/metal/internal/app"
)

// metricsStorage - интерфейс хранилища метрик
type metricsStorage interface {
	// Add - метод добавления/обновления метрики
	Add(ctx context.Context, m app.Metrics) error
	// AddList - метод добавления/обновления списка метрик
	AddList(ctx context.Context, m []app.Metrics) error
	// Get - метод поиска метрики по названию и типу
	Get(ctx context.Context, name string, mType app.MetricType) (app.Metrics, error)
	// List - метод получения списка всех метрик
	List(ctx context.Context) ([]app.Metrics, error)
	// Ping - метод проверки работоспособности хранилища
	Ping(ctx context.Context) error
	// Close - метод закрытия сессии работы хранилища
	Close()
}
