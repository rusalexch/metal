package server

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rusalexch/metal/internal/handlers"
)

func TestNew(t *testing.T) {
	type args struct {
		handler *handlers.Handlers
	}

	h := &handlers.Handlers{}
	tests := []struct {
		want *Server
		args args
		name string
	}{
		{
			name: "should be created server",
			args: args{
				handler: h,
			},
			want: &Server{
				handler: h,
				srv:     http.Server{Addr: "127.0.0.1:8080", Handler: h},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.handler, "127.0.0.1:8080")
			assert.Equal(t, tt.want, got)
		})
	}
}
