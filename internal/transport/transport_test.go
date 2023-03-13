package transport

import (
	"net/http"
	"testing"

	"github.com/rusalexch/metal/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestClient_url(t *testing.T) {
	type fields struct {
		addr   string
		port   int
		client *http.Client
	}
	type args struct {
		m app.Metric
	}

	f := fields{
		addr:   "http://127.0.0.1",
		port:   8080,
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
				m: app.Metric{
					Type:      app.Counter,
					Value:     "123",
					Timestamp: 100,
					Name:      "testCounter",
				},
			},
			want: "http://127.0.0.1:8080/update/counter/testCounter/123",
		},
		{
			name:   "test generation url for guage metric",
			fields: f,
			args: args{
				m: app.Metric{
					Type:      app.Guage,
					Value:     "0.123",
					Timestamp: 100,
					Name:      "testGuage",
				},
			},
			want: "http://127.0.0.1:8080/update/gauge/testGuage/0.123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				addr:   tt.fields.addr,
				port:   tt.fields.port,
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
		port int
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "created client, with addr and port",
			args: args{
				addr: "http://127.0.0.1",
				port: 8080,
			},
			want: &Client{
				addr:   "http://127.0.0.1",
				port:   8080,
				client: &http.Client{},
			},
		},
		{
			name: "created client, without addr",
			args: args{
				addr: "",
				port: 8080,
			},
			want: &Client{
				addr:   "http://127.0.0.1",
				port:   8080,
				client: &http.Client{},
			},
		},
		{
			name: "created client, without port",
			args: args{
				addr: "http://127.0.0.1",
				port: 0,
			},
			want: &Client{
				addr:   "http://127.0.0.1",
				port:   8080,
				client: &http.Client{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.addr, tt.args.port)
			assert.Equal(t, tt.want, got)
		})
	}
}
