package app

func IsMetricType(t MetricType) bool {
	return t == Counter || t == Gauge
}

// guage преобразование метрики типа goage
func ToGauge(v float64, name string) Metrics {
	return Metrics{
		Type:  Gauge,
		Value: &v,
		ID:    name,
	}
}

// counter преобразование метрики типа counter
func ToCounter(v int64, name string) Metrics {
	return Metrics{
		Type:  Counter,
		Delta: &v,
		ID:    name,
	}
}
