package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/rusalexch/metal/internal/hash"
	"github.com/rusalexch/metal/internal/storage"
)

// Handlers структура Хэндлера
type Handlers struct {
	// storage интерфейс хранилища
	storage storage.MetricsStorage
	hash    hash.Hasher
	*chi.Mux
}
