package agent

import "github.com/rusalexch/metal/internal/app"

// Metrics интерфейс для сканера метрик
type Metrics interface {
	// Scan метод для сканирования метрик
	Scan() []app.Metrics
}

// Transport интерфейс клиента для отправки метрик на сервер
type Transport interface {
	// SendOne метод для отправки одной метрики
	SendOne(m app.Metrics) error
	// SendOneJSON метод для отправки одной метрики JSON
	SendOneJSON(m app.Metrics) error
}

// Cache интерфейс кэша для хранения метрик
type Cache interface {
	// Add метод для добавления метрики в кэш
	Add(m []app.Metrics)
	// Reset метод для очистки кэша
	Reset()
	// Get метод для получения кэша
	Get() []app.Metrics
}

