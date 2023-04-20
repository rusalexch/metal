package handlers

import (
	"context"

	"github.com/rusalexch/metal/internal/app"
)

type storager interface {
	Add(ctx context.Context, m app.Metrics) error
	AddList(ctx context.Context, m []app.Metrics) error
	Get(ctx context.Context, name string, mType app.MetricType) (app.Metrics, error)
	List(ctx context.Context) ([]app.Metrics, error)
	Ping(ctx context.Context) error
	Close()
}
