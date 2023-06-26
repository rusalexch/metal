package app

// MetricType тип для типа метрики.
type MetricType = string

// Metrics структура метрики.
type Metrics struct {
	// ID наименование метрики.
	ID string `json:"id"`
	// Type наименование типа метрики.
	Type MetricType `json:"type"`
	// Delta значение метрики типа counter.
	Delta *int64 `json:"delta,omitempty"`
	// Value значение метрики типа gauge.
	Value *float64 `json:"value,omitempty"`
	// Hash значение хеш-функции.
	Hash string `json:"hash,omitempty"`
}
