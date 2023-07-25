package server

import (
	"net/http"

	"github.com/rusalexch/metal/internal/handlers"
)

// Server - структура сервера.
type Server struct {
	// handler - указатель на Хендлеры.
	handler *handlers.Handlers
	// srv - инстанс сервера
	srv http.Server
}
