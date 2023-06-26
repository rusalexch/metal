package app

func IsMetricType(t MetricType) bool {
	return t == Counter || t == Gauge
}

// AsGauge - получение метрики типа goage.
func AsGauge(v float64, name string) Metrics {
	return Metrics{
		Type:  Gauge,
		Value: &v,
		ID:    name,
	}
}

// AsCounter - получение метрики типа counter.
func AsCounter(v int64, name string) Metrics {
	return Metrics{
		Type:  Counter,
		Delta: &v,
		ID:    name,
	}
}
