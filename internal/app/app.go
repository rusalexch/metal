package app

func IsMetricType(t MetricType) bool {
	return t == Counter || t == Gauge
}
