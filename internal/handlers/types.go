package handlers

import (
	"time"

	"github.com/go-chi/chi/v5"
)

// Handlers - структура Хэндлера.
type Handlers struct {
	// storage - хранилище.
	storage storager
	// hash - хэш-функция.
	hash hasher
	// timeout - интервал таймаута.
	timeout time.Duration
	*chi.Mux
}
