package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/services"
	"github.com/rusalexch/metal/internal/storage"
	"github.com/rusalexch/metal/internal/utils"
)

func (h *Handlers) find(w http.ResponseWriter, r *http.Request) {
	fmt.Println("value")
	ID := chi.URLParam(r, "ID")
	mType := chi.URLParam(r, "mType")
	m, err := h.services.MetricsService.Get(ID, mType)
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
	switch m.Type {
	case app.Counter:
		fmt.Fprint(w, utils.Int64ToStr(*m.Delta))
	case app.Gauge:
		fmt.Fprint(w, utils.Float64ToStr(*m.Value))
	}
}
