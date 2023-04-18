package storage

import "github.com/rusalexch/metal/internal/app"

// MetricsStorage интерфейс хранилища метрик
type MetricsStorage interface {
	Add(m app.Metrics) error
	AddList(m []app.Metrics) error
	Get(name string, mType app.MetricType) (app.Metrics, error)
	List() ([]app.Metrics, error)
	Ping() error
	Close()
}
