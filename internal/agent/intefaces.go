package agent

// Metrics интерфейс для сканера метрик
type Metrics interface {
	// Scan метод для сканирования метрик
	Scan() ([]Metric, error)
}

// Transport интерфейс клиента для отправки метрик на сервер
type Transport interface {
	// SendOne метод для отправки одной метрики
	SendOne(m Metric) error
}

// Cache интерфейс кэша для хранения метрик
type Cache interface {
	// Add метод для добавления метрики в кэш
	Add(m []Metric) error
	// Reset метод для очистки кэша
	Reset() error
	// Get метод для получения кэша
	Get() ([]Metric, error)
}
