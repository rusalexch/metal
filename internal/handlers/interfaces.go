package handlers

import "github.com/rusalexch/metal/internal/app"

type storager interface {
	Add(m app.Metrics) error
	Get(name string, mType app.MetricType) (app.Metrics, error)
	List() ([]app.Metrics, error)
	Ping() error
	Close()
}
