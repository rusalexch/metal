package services

import "github.com/rusalexch/metal/internal/app"

// Mertrics интерфейс сервиса метрик
type Mertrics interface {
	Add(m app.Metrics) error
	Get(name string, mType app.MetricType) (app.Metrics, error)
	List() []app.Metrics
	Subscribe(f func())
}

type Healther interface {
	Ping() error
}

type Repositorier interface {
	Ping() error
}
