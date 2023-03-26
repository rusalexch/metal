package server

import (
	"github.com/rusalexch/metal/internal/handlers"
)

// Server структура сервера
type Server struct {
	// addr адрес сервера
	baseURL string
	// port порт сервера
	port int
	// handler указатель на Хендлеры
	handler *handlers.Handlers
}
