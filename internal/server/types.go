package server

import (
	"github.com/rusalexch/metal/internal/handlers"
)

// Server - структура сервера.
type Server struct {
	// handler - указатель на Хендлеры.
	handler *handlers.Handlers
	// addr - адрес сервера.
	addr string
}
