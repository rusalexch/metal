package agent

import "github.com/rusalexch/metal/internal/app"

// Metrics интерфейс для сканера метрик
type Metrics interface {
	// Scan метод для сканирования метрик
	Scan() []app.Metric
}

// Transport интерфейс клиента для отправки метрик на сервер
type Transport interface {
	// SendOne метод для отправки одной метрики
	SendOne(m app.Metric) error
}

// Cache интерфейс кэша для хранения метрик
type Cache interface {
	// Add метод для добавления метрики в кэш
	Add(m []app.Metric) error
	// Reset метод для очистки кэша
	Reset() error
	// Get метод для получения кэша
	Get() ([]app.Metric, error)
}
