package handlers

import (
	"log"
	"net/http"
)

// ping хендлер для проверки работоспособности
func (h *Handlers) ping(w http.ResponseWriter, r *http.Request) {
	if err := h.services.HealthCheck.Ping(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
