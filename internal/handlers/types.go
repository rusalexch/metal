package handlers

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rusalexch/metal/internal/hash"
)

// Handlers структура Хэндлера
type Handlers struct {
	// storage интерфейс хранилища
	storage storager
	hash    hash.Hasher
	timeout time.Duration
	*chi.Mux
}
