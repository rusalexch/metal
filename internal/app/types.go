package app

// MetricType тип для типа метрики
type MetricType = string

// Metric структура метрики
type Metric struct {
	// ID наименование метрики
	ID string
	// Type наименование типа метрики
	Type MetricType
	// Value значение метрики в строковом формате
	Value string
}
