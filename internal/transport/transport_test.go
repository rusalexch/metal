package transport

import (
	"net/http"
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

func TestClient_url(t *testing.T) {
	type fields struct {
		addr   string
		client *http.Client
	}
	type args struct {
		m app.Metrics
	}

	f := fields{
		addr:   "127.0.0.1:8080",
		client: &http.Client{},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "test generation url for counter metric",
			fields: f,
			args: args{
				m: app.Metrics{
					Type:  app.Counter,
					Delta: int64AsPointer(123),
					ID:    "testCounter",
				},
			},
			want: "http://127.0.0.1:8080/update/counter/testCounter/123",
		},
		{
			name:   "test generation url for guage metric",
			fields: f,
			args: args{
				m: app.Metrics{
					Type:  app.Gauge,
					Value: float64AsPointer(0.123),
					ID:    "testGuage",
				},
			},
			want: "http://127.0.0.1:8080/update/gauge/testGuage/0.123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				addr:   tt.fields.addr,
				client: tt.fields.client,
			}
			if got := c.url(tt.args.m); got != tt.want {
				t.Errorf("Client.url() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "created client, with addr and port",
			args: args{
				addr: "127.0.0.1:8080",
			},
			want: &Client{
				addr:   "127.0.0.1:8080",
				client: &http.Client{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.addr)
			assert.Equal(t, tt.want, got)
		})
	}
}
