package server

import (
	"context"
	"net/http"

	"github.com/rusalexch/metal/internal/handlers"
)

// New конструктор сервера сбора метрик
func New(handler *handlers.Handlers) *Server {
	return &Server{
		server: http.Server{
			Addr: "127.0.0.1:8080",
		},
		handler: handler,
	}
}

// Start запуск сервера сбора метрик
func (s *Server) Start() error {
	s.handler.Init()
	err := s.server.ListenAndServe()

	return err
}

// Stop остановка сервера
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
