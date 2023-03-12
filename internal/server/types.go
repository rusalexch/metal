package server

import (
	"net/http"

	"github.com/rusalexch/metal/internal/handlers"
)

// Server структура сервера
type Server struct {
	// server экземпляр сервера
	server http.Server
	// handler указатель на Хендлеры
	handler *handlers.Handlers
}
