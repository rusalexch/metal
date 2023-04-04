package server

import (
	"github.com/rusalexch/metal/internal/handlers"
)

// Server структура сервера
type Server struct {
	// addr адрес сервера
	addr string
	// handler указатель на Хендлеры
	handler *handlers.Handlers
}
