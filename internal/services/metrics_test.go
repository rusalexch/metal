package services

import (
	"testing"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewMertricsService(t *testing.T) {
	type args struct {
		storage storage.MetricsStorage
	}

	s := storage.New()
	tests := []struct {
		name string
		args args
		want *MertricsService
	}{
		{
			name: "created mertics service",
			args: args{
				storage: s,
			},
			want: &MertricsService{
				storage: s,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMertricsService(tt.args.storage)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMertricsService_Add(t *testing.T) {
	type fields struct {
		storage storage.MetricsStorage
	}

	f := fields{
		storage: storage.New(),
	}
	tests := []struct {
		name    string
		fields  fields
		args    app.Metric
		wantErr bool
	}{
		{
			name:   "add new counter metric",
			fields: f,
			args: app.Metric{
				Type:      app.Counter,
				Value:     "777",
				Timestamp: 0,
				Name:      "testCounter1",
			},
			wantErr: false,
		},
		{
			name:   "add new guage metric",
			fields: f,
			args: app.Metric{
				Type:      app.Guage,
				Value:     "0.000002",
				Timestamp: 0,
				Name:      "testGuage2",
			},
			wantErr: false,
		},
		{
			name:   "fault with wrong type",
			fields: f,
			args: app.Metric{
				Type:      "wrongType",
				Value:     "123",
				Timestamp: 0,
				Name:      "test3",
			},
			wantErr: true,
		},
		{
			name:   "fault with wrong counter value",
			fields: f,
			args: app.Metric{
				Type:      app.Counter,
				Value:     "wer",
				Timestamp: 0,
				Name:      "testCounter4",
			},
			wantErr: true,
		},
		{
			name:   "fault with wrong guage value",
			fields: f,
			args: app.Metric{
				Type:      app.Guage,
				Value:     "swq",
				Timestamp: 0,
				Name:      "testCounter5",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &MertricsService{
				storage: tt.fields.storage,
			}
			if err := ms.Add(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("MertricsService.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMertricsService_Get(t *testing.T) {
	type fields struct {
		storage storage.MetricsStorage
	}
	type args struct {
		name  string
		mType app.MetricType
	}
	type want struct {
		isErr bool
		m     app.Metric
	}

	f := fields{
		storage: storage.New(),
	}
	f.storage.AddCounter("testCounter1", 777)
	f.storage.AddCounter("testCounter2", 93245)
	f.storage.AddCounter("testCounter3", -10005)
	f.storage.AddGauge("testGuage1", 0.00001)
	f.storage.AddGauge("testGuage2", 5.3)
	f.storage.AddGauge("testGuage3", -0.000000001)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name:   "get counter metrics",
			fields: f,
			args: args{
				name:  "testCounter3",
				mType: app.Counter,
			},
			want: want{
				isErr: false,
				m: app.Metric{
					Type:      app.Counter,
					Value:     "-10005",
					Timestamp: 0,
					Name:      "testCounter3",
				},
			},
		},
		{
			name:   "get guage metrics",
			fields: f,
			args: args{
				name:  "testGuage2",
				mType: app.Guage,
			},
			want: want{
				isErr: false,
				m: app.Metric{
					Type:      app.Guage,
					Value:     "5.3",
					Timestamp: 0,
					Name:      "testGuage2",
				},
			},
		},
		{
			name:   "fault, not found counter metric",
			fields: f,
			args: args{
				name:  "testCounter333",
				mType: app.Counter,
			},
			want: want{
				isErr: true,
				m:     app.Metric{},
			},
		},
		{
			name:   "fault, not found guage metric",
			fields: f,
			args: args{
				name:  "testGuage23",
				mType: app.Guage,
			},
			want: want{
				isErr: true,
				m:     app.Metric{},
			},
		},
		{
			name:   "fault with error type",
			fields: f,
			args: args{
				name:  "testGuage2",
				mType: "wrongType",
			},
			want: want{
				isErr: true,
				m:     app.Metric{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &MertricsService{
				storage: tt.fields.storage,
			}
			got, err := ms.Get(tt.args.name, tt.args.mType)
			if (err != nil) != tt.want.isErr {
				t.Errorf("MertricsService.Get() error = %v, wantErr %v", err, tt.want.isErr)
				return
			}
			assert.Equal(t, tt.want.m, got)
		})
	}
}
