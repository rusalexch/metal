package services

import (
	"testing"

	"github.com/rusalexch/metal/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		storage storage.MetricsStorage
	}
	s := storage.New(nil)
	ms := NewMertricsService(s)
	hs := NewHealthCheck(s)
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
				HealthCheck: hs,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.storage)
			assert.Equal(t, tt.want, got)
		})
	}
}
