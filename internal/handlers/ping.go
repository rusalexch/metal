package handlers

import (
	"context"
	"log"
	"net/http"
)

// ping хендлер для проверки работоспособности
func (h *Handlers) ping(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	if err := h.storage.Ping(ctx); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
