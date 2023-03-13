package agent

import (
	"reflect"
	"testing"
	"time"

	"github.com/rusalexch/metal/internal/cashe"
	"github.com/rusalexch/metal/internal/metric"
	"github.com/rusalexch/metal/internal/transport"
)

func TestNew(t *testing.T) {
	type args struct {
		conf Config
	}
	m := metric.New()
	c := cashe.New()
	tr := transport.New("http://127.0.0.1", 8080)
	tests := []struct {
		name string
		args args
		want *App
	}{
		{
			name: "should be created",
			args: args{
				conf: Config{
					PollInterval:   2 * time.Second,
					ReportInterval: 10 * time.Second,
					Metrics:        m,
					Cache:          c,
					Transport:      tr,
				},
			},
			want: &App{
				pollInterval:   2 * time.Second,
				reportInterval: 10 * time.Second,
				metrics:        m,
				cache:          c,
				transport:      tr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

