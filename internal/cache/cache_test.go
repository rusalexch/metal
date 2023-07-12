package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/utils"
)

func TestNew(t *testing.T) {
	tests := []struct {
		want *Cache
		name string
	}{
		{
			name: "should be created",
			want: &Cache{
				m: map[string]app.Metrics{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCache_add(t *testing.T) {
	type args struct {
		m app.Metrics
	}
	type want struct {
		id    string
		delta int64
		value float64
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "add counter",
			args: args{
				m: app.Metrics{
					ID:    "testCounter1",
					Type:  app.Counter,
					Delta: utils.Int64AsPointer(12),
				},
			},
			want: want{
				id:    "testCounter1",
				delta: 12,
			},
		},
		{
			name: "add gauge",
			args: args{
				m: app.Metrics{
					ID:    "testGauge1",
					Type:  app.Gauge,
					Value: utils.Float64AsPointer(0.00001),
				},
			},
			want: want{
				id:    "testGauge1",
				value: 0.00001,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			c.add(tt.args.m)
			got, ok := c.m[tt.want.id]
			assert.Equal(t, true, ok)
			if tt.args.m.Type == app.Counter {
				assert.Equal(t, tt.want.delta, *got.Delta)
			} else if tt.args.m.Type == app.Gauge {
				assert.Equal(t, tt.want.value, *got.Value)
			}
		})
	}
}

func TestCache_get(t *testing.T) {
	type args struct {
		m []app.Metrics
	}
	type want struct {
		counter app.Metrics
		gauge   app.Metrics
	}

	mCounter := app.Metrics{
		ID:    "TestCounter1",
		Type:  app.Counter,
		Delta: utils.Int64AsPointer(15),
	}
	mGauge := app.Metrics{
		ID:    "TestGauge1",
		Type:  app.Gauge,
		Value: utils.Float64AsPointer(0.001),
	}

	tests := []struct {
		name string
		want want
		args args
	}{
		{
			name: "get from cache",
			args: args{
				m: []app.Metrics{mCounter, mGauge},
			},
			want: want{
				counter: mCounter,
				gauge:   mGauge,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			for _, m := range tt.args.m {
				c.add(m)
			}
			got := c.get()
			assert.Contains(t, got, tt.want.counter)
			assert.Contains(t, got, tt.want.gauge)
		})
	}
}
