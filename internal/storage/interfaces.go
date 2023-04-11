package storage

// MetricsStorage интерфейс хранилища метрик
type MetricsStorage interface {
	AddCounter(name string, value int64) error
	AddGauge(name string, value float64) error
	GetCounter(name string) (int64, error)
	GetGauge(name string) (float64, error)
	ListCounter() []MetricCounter
	ListGauge() []MetricGauge
	Ping() error
}
