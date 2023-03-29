package server

import (
	"net/http"

	"github.com/rusalexch/metal/internal/handlers"
)

// New конструктор сервера сбора метрик
func New(handler *handlers.Handlers, addr string) *Server {
	return &Server{
		addr:    addr,
		handler: handler,
	}
}

// Start запуск сервера сбора метрик
func (s *Server) Start() error {
	s.handler.Init()

	err := http.ListenAndServe(s.addr, s.handler.Mux)

	return err
}
