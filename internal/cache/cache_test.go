package cache

import (
	"testing"

	"github.com/rusalexch/metal/internal/app"
	"github.com/stretchr/testify/assert"
)

func int64AsPointer(v int64) *int64 {
	return &v
}

func float64AsPointer(v float64) *float64 {
	return &v
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Cache
	}{
		{
			name: "should be created",
			want: &Cache{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCache_Add(t *testing.T) {
	type want struct {
		m1 []app.Metrics
		m2 []app.Metrics
	}
	type args struct {
		firstAdd  []app.Metrics
		secondAdd []app.Metrics
	}

	metricsFirst := []app.Metrics{
		{
			Type:  app.Counter,
			Delta: int64AsPointer(64),
			ID:    "testCounter1",
		},
		{
			Type:  app.Gauge,
			Value: float64AsPointer(1.23),
			ID:    "testGuage1",
		},
		{
			Type:  app.Gauge,
			Value: float64AsPointer(-1.000001),
			ID:    "testGuage2",
		},
	}
	metricsSecond := []app.Metrics{
		{
			Type:  app.Counter,
			Delta: int64AsPointer(1230),
			ID:    "testCounter5",
		},
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "added three metrics",
			args: args{
				firstAdd:  metricsFirst,
				secondAdd: metricsSecond,
			},
			want: want{
				m1: metricsFirst,
				m2: append(metricsFirst, metricsSecond...),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			c.Reset()
			c.Add(tt.args.firstAdd)

			assert.Equal(t, tt.want.m1, c.m)

			c.Add(tt.args.secondAdd)
			assert.Equal(t, tt.want.m2, c.m)
		})
	}
}

func TestCache_Reset(t *testing.T) {
	type want struct {
		m []app.Metrics
	}
	metrics := []app.Metrics{
		{
			Type:  app.Counter,
			Delta: int64AsPointer(123),
			ID:    "testCounter1",
		},
		{
			Type:  app.Gauge,
			Value: float64AsPointer(1.23),
			ID:    "testGuage1",
		},
		{
			Type:  app.Gauge,
			Value: float64AsPointer(-1.000001),
			ID:    "testGuage2",
		},
	}

	tests := []struct {
		name   string
		fields []app.Metrics
		want   want
	}{
		{
			name:   "reset when empty",
			fields: []app.Metrics{},
			want: want{
				m: []app.Metrics{},
			},
		},
		{
			name:   "reset `when exist",
			fields: metrics,
			want: want{
				m: []app.Metrics{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			if len(tt.fields) != 0 {
				c.Add(tt.fields)
			}
			c.Reset()
			assert.Equal(t, tt.want.m, c.m)
		})
	}
}

func TestCache_Get(t *testing.T) {
	type want struct {
		m []app.Metrics
	}
	metrics := []app.Metrics{
		{
			Type:  app.Counter,
			Delta: int64AsPointer(123),
			ID:    "testCounter1",
		},
		{
			Type:  app.Gauge,
			Value: float64AsPointer(1.23),
			ID:    "testGuage1",
		},
		{
			Type:  app.Gauge,
			Value: float64AsPointer(-1.000001),
			ID:    "testGuage2",
		},
	}
	tests := []struct {
		name   string
		fields []app.Metrics
		want   want
	}{
		{
			name:   "get empty cache",
			fields: []app.Metrics{},
			want: want{
				m: []app.Metrics{},
			},
		},
		{
			name:   "get values from cache",
			fields: metrics,
			want: want{
				m: metrics,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			c.Reset()
			if len(tt.fields) != 0 {
				c.Add(tt.fields)
			}
			got := c.Get()
			assert.Equal(t, tt.want.m, got)
		})
	}
}
