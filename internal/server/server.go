package server

import (
	"context"
	"log"
	"net/http"

	"github.com/rusalexch/metal/internal/handlers"
)

// New - конструктор сервера сбора метрик.
func New(handler *handlers.Handlers, addr string) *Server {
	return &Server{
		handler: handler,
		srv:     http.Server{Addr: addr, Handler: handler},
	}
}

// Start - запуск сервера сбора метрик.
func (s *Server) Start() {
	s.handler.Init()
	err := s.srv.ListenAndServe()
	log.Println(err)
}

// Shutdown - остановка сервера
func (s *Server) Shutdown(ctx context.Context, ch chan struct{}) {
	err := s.srv.Shutdown(ctx)
	if err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}
	close(ch)
}
