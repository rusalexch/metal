package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/rusalexch/metal/internal/services"
)

// Handlers структура Хэндлера
type Handlers struct {
	// services указатель на сервисы
	services *services.Services
	*chi.Mux
}
