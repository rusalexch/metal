package handlers

import (
	"time"

	"github.com/go-chi/chi/v5"
)

// Handlers структура Хэндлера
type Handlers struct {
	// storage интерфейс хранилища
	storage storager
	hash    hasher
	timeout time.Duration
	*chi.Mux
}
