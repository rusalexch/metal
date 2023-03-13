package server

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/rusalexch/metal/internal/handlers"
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
				server: http.Server{
					Addr: "127.0.0.1:8080",
				},
				handler: h,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
