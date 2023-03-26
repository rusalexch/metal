package server

import (
	"testing"

	"github.com/rusalexch/metal/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		handler *handlers.Handlers
	}

	h := &handlers.Handlers{}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "should be created server",
			args: args{
				handler: h,
			},
			want: &Server{
				baseURL: "127.0.0.1",
				port:    8080,
				handler: h,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.handler, "", 0)
			assert.Equal(t, tt.want, got)
		})
	}
}
