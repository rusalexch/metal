package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Storage
	}{
		{
			name: "created storage",
			want: &Storage{
				counters: map[string]MetricCounter{},
				gauges:   map[string]MetricGauge{},
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

func TestStorage_AddCounter(t *testing.T) {
	type storage struct {
		counters map[string]MetricCounter
		gauges   map[string]MetricGauge
	}
	type args struct {
		name  string
		value int64
	}
	type want struct {
		isErr   bool
		metrics []MetricCounter
	}

	st := storage{
		counters: map[string]MetricCounter{},
		gauges:   map[string]MetricGauge{},
	}
	tests := []struct {
		name    string
		storage storage
		args    []args
		want    want
	}{
		{
			name:    "added one counter metric",
			storage: st,
			args: []args{
				{
					name:  "testCounter",
					value: 123,
				},
			},
			want: want{
				isErr: false,
				metrics: []MetricCounter{
					{
						Value: 123,
						Name:  "testCounter",
					},
				},
			},
		},
		{
			name:    "added three counter metrics",
			storage: st,
			args: []args{
				{
					name:  "testCounter1",
					value: 1,
				},
				{
					name:  "testCounter2",
					value: 2,
				},
				{
					name:  "testCounter3",
					value: 3,
				},
			},
			want: want{
				isErr: false,
				metrics: []MetricCounter{
					{
						Value: 1,
						Name:  "testCounter1",
					},
					{
						Value: 2,
						Name:  "testCounter2",
					},
					{
						Value: 3,
						Name:  "testCounter3",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				counters: tt.storage.counters,
				gauges:   tt.storage.gauges,
			}
			for _, v := range tt.args {
				if err := s.AddCounter(v.name, v.value); (err != nil) != tt.want.isErr {
					t.Fatalf("Storage.AddCounter() error = %v, wantErr %v", err, tt.want.isErr)
				}
			}

			got := make([]MetricCounter, len(tt.want.metrics))
			for i, m := range tt.want.metrics {
				got[i] = s.counters[m.Name]
			}
			assert.Equal(t, tt.want.metrics, got)
		})
	}
}

func TestStorage_AddGuage(t *testing.T) {
	type fields struct {
		counters map[string]MetricCounter
		gauges   map[string]MetricGauge
	}
	type args struct {
		name  string
		value float64
	}
	type want struct {
		isErr   bool
		metrics []MetricGauge
	}

	f := fields{
		counters: map[string]MetricCounter{},
		gauges:   map[string]MetricGauge{},
	}
	tests := []struct {
		name   string
		fields fields
		args   []args
		want   want
	}{
		{
			name:   "aded one gauge metric",
			fields: f,
			args: []args{
				{
					name:  "testGuage",
					value: 3.14,
				},
			},
			want: want{
				isErr: false,
				metrics: []MetricGauge{
					{
						Value: 3.14,
						Name:  "testGuage",
					},
				},
			},
		},
		{
			name:   "aded three guage metrics",
			fields: f,
			args: []args{
				{
					name:  "testGuage1",
					value: 3.14,
				},
				{
					name:  "testGuage2",
					value: 45.14,
				},
				{
					name:  "testGuage3",
					value: -0.0000000001,
				},
			},
			want: want{
				isErr: false,
				metrics: []MetricGauge{
					{
						Value: 3.14,
						Name:  "testGuage1",
					},
					{
						Value: 45.14,
						Name:  "testGuage2",
					},
					{
						Value: -0.0000000001,
						Name:  "testGuage3",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				counters: tt.fields.counters,
				gauges:   tt.fields.gauges,
			}
			for _, arg := range tt.args {
				if err := s.AddGauge(arg.name, arg.value); (err != nil) != tt.want.isErr {
					t.Fatalf("Storage.AddGuage() error = %v, wantErr %v", err, tt.want.isErr)
				}
			}

			got := make([]MetricGauge, len(tt.want.metrics))
			for i, m := range tt.want.metrics {
				got[i] = s.gauges[m.Name]
			}
			assert.Equal(t, tt.want.metrics, got)
		})
	}
}

func TestStorage_GetCounter(t *testing.T) {
	type fields struct {
		counters map[string]MetricCounter
		gauges   map[string]MetricGauge
	}
	type args struct {
		name string
	}
	type want struct {
		isErr bool
		value int64
	}

	f := fields{
		counters: map[string]MetricCounter{
			"testCounter1": {Name: "testCounter1", Value: 1001},
			"testCounter2": {Name: "testCounter2", Value: 1002},
			"testCounter3": {Name: "testCounter3", Value: 1003},
		},
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name:   "should be found value",
			fields: f,
			args: args{
				name: "testCounter2",
			},
			want: want{
				isErr: false,
				value: 1002,
			},
		},
		{
			name:   "should be not found value",
			fields: f,
			args: args{
				name: "testCounter4",
			},
			want: want{
				isErr: true,
				value: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				counters: tt.fields.counters,
				gauges:   tt.fields.gauges,
			}
			got, err := s.GetCounter(tt.args.name)
			if (err != nil) != tt.want.isErr {
				t.Errorf("Storage.GetCounter() error = %v, wantErr %v", err, tt.want.isErr)
				return
			}
			if got != tt.want.value {
				t.Errorf("Storage.GetCounter() = %v, want %v", got, tt.want.value)
			}
		})
	}
}

func TestStorage_GetGuage(t *testing.T) {
	type fields struct {
		counters map[string]MetricCounter
		gauges   map[string]MetricGauge
	}
	type args struct {
		name string
	}
	type want struct {
		isErr bool
		value float64
	}

	f := fields{
		counters: map[string]MetricCounter{},
		gauges: map[string]MetricGauge{
			"testGuage1": {Name: "testGuage1", Value: 3.14},
			"testGuage2": {Name: "testGuage2", Value: -3.14},
			"testGuage3": {Name: "testGuage3", Value: 0.000000001},
		},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name:   "should found value",
			fields: f,
			args: args{
				name: "testGuage3",
			},
			want: want{
				isErr: false,
				value: 0.000000001,
			},
		},
		{
			name:   "should not found value",
			fields: f,
			args: args{
				name: "testGuage30",
			},
			want: want{
				isErr: true,
				value: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				counters: tt.fields.counters,
				gauges:   tt.fields.gauges,
			}
			got, err := s.GetGauge(tt.args.name)
			if (err != nil) != tt.want.isErr {
				t.Errorf("Storage.GetGuage() error = %v, wantErr %v", err, tt.want.isErr)
				return
			}
			if got != tt.want.value {
				t.Errorf("Storage.GetGuage() = %v, want %v", got, tt.want.value)
			}
		})
	}
}
