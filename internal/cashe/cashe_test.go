package cashe

import (
	"reflect"
	"testing"

	"github.com/rusalexch/metal/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Cashe
	}{
		{
			name: "should be created",
			want: &Cashe{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCashe_Add(t *testing.T) {
	type want struct {
		isErr bool
		m1    []app.Metric
		m2    []app.Metric
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
			Value:     "123",
			Timestamp: 1,
			Name:      "testCounter1",
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
				isErr: false,
				m1:    metricsFirst,
				m2:    append(metricsFirst, metricsSecond...),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			if err := c.Add(tt.args.firstAdd); (err != nil) != tt.want.isErr {
				t.Errorf("Cashe.Add() error = %v, wantErr %v", err, tt.want.isErr)
				return
			}
			assert.Equal(t, tt.want.m1, c.m)
			if err := c.Add(tt.args.secondAdd); (err != nil) != tt.want.isErr {
				t.Errorf("Cashe.Add() error = %v, wantErr %v", err, tt.want.isErr)
				return
			}
			assert.Equal(t, tt.want.m2, c.m)
		})
	}
}

func TestCashe_Reset(t *testing.T) {
	type want struct {
		isErr bool
		m     []app.Metric
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
				isErr: false,
				m:     []app.Metric{},
			},
		},
		{
			name:   "reset `when exist",
			fields: metrics,
			want: want{
				isErr: false,
				m:     []app.Metric{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			if len(tt.fields) != 0 {
				c.Add(tt.fields)
			}
			if err := c.Reset(); (err != nil) != tt.want.isErr {
				t.Errorf("Cashe.Reset() error = %v, wantErr %v", err, tt.want.isErr)
			}
			assert.Equal(t, tt.want.m, c.m)
		})
	}
}

func TestCashe_Get(t *testing.T) {
	type want struct {
		isErr bool
		m     []app.Metric
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
				isErr: false,
				m:     []app.Metric{},
			},
		},
		{
			name:   "get values from cache",
			fields: metrics,
			want: want{
				isErr: false,
				m:     metrics,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			if len(tt.fields) != 0 {
				c.Add(tt.fields)
			}
			got, err := c.Get()
			if (err != nil) != tt.want.isErr {
				t.Errorf("Cashe.Get() error = %v, wantErr %v", err, tt.want.isErr)
				return
			}
			assert.Equal(t, tt.want.m, got)
		})
	}
}