package storage

// MetricsStorage интерфейс хранилища метрик
type MetricsStorage interface {
	AddCounter(name string, value int64) error
	AddGuage(name string, value float64) error
	GetCounter(name string) (int64, error)
	GetGuage(name string) (float64, error)
}
