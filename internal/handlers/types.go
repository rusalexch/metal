package handlers

import (
	"crypto/rsa"
	"time"

	"github.com/go-chi/chi/v5"
)

// Handlers - структура Хэндлера.
type Handlers struct {
	*chi.Mux
	// storage - хранилище.
	storage storager
	// hash - хэш-функция.
	hash hasher
	// timeout - интервал таймаута.
	timeout time.Duration
	// privateKey - приватный ключ
	privateKey *rsa.PrivateKey
}
