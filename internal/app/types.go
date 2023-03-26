package app

// MetricType тип для типа метрики
type MetricType = string

// Metric структура метрики
type Metric struct {
	// Type наименование типа метрики
	Type MetricType
	// Value значение метрики в строковом формате
	Value string
	// Name наименование метрики
	Name string
}
