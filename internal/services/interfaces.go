package services

import "github.com/rusalexch/metal/internal/app"

// Mertrics интерфейс сервиса метрик
type Mertrics interface {
	Add(m app.Metric) error
	Get(name string, mType app.MetricType) (app.Metric, error)
	List() []app.Metric
}
