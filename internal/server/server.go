package server

import (
	"fmt"
	"net/http"

	"github.com/rusalexch/metal/internal/handlers"
)

// New конструктор сервера сбора метрик
func New(handler *handlers.Handlers, baseURL string, port int) *Server {
	if baseURL == "" {
		baseURL = "127.0.0.1"
	}
	if port == 0 {
		port = 8080
	}
	return &Server{
		baseURL: baseURL,
		port:    port,
		handler: handler,
	}
}

// Start запуск сервера сбора метрик
func (s *Server) Start() error {
	s.handler.Init()

	err := http.ListenAndServe(s.addr(), s.handler.Mux)

	return err
}

// addr метод получение адреса сервера
func (s *Server) addr() string {
	return fmt.Sprintf("%s:%d", s.baseURL, s.port)
}
