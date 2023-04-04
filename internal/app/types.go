package app

// MetricType тип для типа метрики
type MetricType = string

// Metrics структура метрики
type Metrics struct {
	// ID наименование метрики
	ID string `json:"id"`
	// Type наименование типа метрики
	Type MetricType `json:"type"`
	// Value значение метрики в случае передачи counter
	Delta *int64 `json:"delta,omitempty"`
	// Value значение метрики в случае передачи gauge
	Value *float64 `json:"value,omitempty"`
}
