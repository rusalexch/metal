package cache

import (
	"testing"

	"github.com/rusalexch/metal/internal/app"
	"github.com/stretchr/testify/assert"
)

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
		m1 []app.Metric
		m2 []app.Metric
	}
	type args struct {
		firstAdd  []app.Metric
		secondAdd []app.Metric
	}

	metricsFirst := []app.Metric{
		{
			Type:      app.Counter,
			Value:     "123",
			Timestamp: 0,
			Name:      "testCounter1",
		},
		{
			Type:      app.Guage,
			Value:     "1.23",
			Timestamp: 0,
			Name:      "testGuage1",
		},
		{
			Type:      app.Guage,
			Value:     "-1.000001",
			Timestamp: 0,
			Name:      "testGuage2",
		},
	}
	metricsSecond := []app.Metric{
		{
			Type:      app.Counter,
			Value:     "1230",
			Timestamp: 1,
			Name:      "testCounter5",
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
		m []app.Metric
	}
	metrics := []app.Metric{
		{
			Type:      app.Counter,
			Value:     "123",
			Timestamp: 0,
			Name:      "testCounter1",
		},
		{
			Type:      app.Guage,
			Value:     "1.23",
			Timestamp: 0,
			Name:      "testGuage1",
		},
		{
			Type:      app.Guage,
			Value:     "-1.000001",
			Timestamp: 0,
			Name:      "testGuage2",
		},
	}

	tests := []struct {
		name   string
		fields []app.Metric
		want   want
	}{
		{
			name:   "reset when empty",
			fields: []app.Metric{},
			want: want{
				m: []app.Metric{},
			},
		},
		{
			name:   "reset `when exist",
			fields: metrics,
			want: want{
				m: []app.Metric{},
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
		m []app.Metric
	}
	metrics := []app.Metric{
		{
			Type:      app.Counter,
			Value:     "123",
			Timestamp: 0,
			Name:      "testCounter1",
		},
		{
			Type:      app.Guage,
			Value:     "1.23",
			Timestamp: 0,
			Name:      "testGuage1",
		},
		{
			Type:      app.Guage,
			Value:     "-1.000001",
			Timestamp: 0,
			Name:      "testGuage2",
		},
	}
	tests := []struct {
		name   string
		fields []app.Metric
		want   want
	}{
		{
			name:   "get empty cache",
			fields: []app.Metric{},
			want: want{
				m: []app.Metric{},
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
