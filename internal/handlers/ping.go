package handlers

import (
	"log"
	"net/http"
)

// ping хендлер для проверки работоспособности
func (h *Handlers) ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.storage.Ping(ctx); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
