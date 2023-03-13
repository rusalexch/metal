package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rusalexch/metal/internal/services"
	"github.com/rusalexch/metal/internal/storage"
)

func (h *Handlers) find(w http.ResponseWriter, r *http.Request) {
	fmt.Println("value")
	name := chi.URLParam(r, "name")
	mType := chi.URLParam(r, "mType")
	m, err := h.services.MetricsService.Get(name, mType)
	if err != nil {
		if errors.Is(err, services.ErrIncorrectType) {
			w.WriteHeader(http.StatusNotImplemented)
			fmt.Fprint(w, "method not implemented")
			return
		}
		if errors.Is(err, storage.ErrMetricNotFound) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, err)
			return
		}
	}
	fmt.Fprint(w, m.Value)
}
