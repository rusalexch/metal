package services

import (
	"reflect"
	"testing"

	"github.com/rusalexch/metal/internal/storage"
)

func TestNew(t *testing.T) {
	type args struct {
		storage storage.MetricsStorage
	}
	s := storage.New()
	ms := NewMertricsService(s)
	tests := []struct {
		name string
		args args
		want *Services
	}{
		{
			name: "should created services",
			args: args{
				storage: s,
			},
			want: &Services{
				MetricsService: ms,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
