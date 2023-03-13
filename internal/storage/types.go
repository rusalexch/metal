package storage

// MetricCounter структура метрики counter для хранения
type MetricCounter struct {
	// Value значение метрики в строковом формате
	Value int64
	// Name наименование метрики
	Name string
}

// MetricGuage структура метрики guage для хранения
type MetricGauge struct {
	// Value значение метрики в строковом формате
	Value float64
	// Name наименование метрики
	Name string
}

// Storage структура хранилища
type Storage struct {
	// counters мапа хранения для метрик типа counter
	counters map[string]MetricCounter
	// guages апа хранения для метрик типа guage
	gauges map[string]MetricGauge
}
