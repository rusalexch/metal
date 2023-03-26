package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/services"
)

// update Хэндлер для обновления метрик
func (h *Handlers) update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update")

	m := app.Metric{
		Type:      chi.URLParam(r, "mType"),
		ID:      	 chi.URLParam(r, "ID"),
		Value:     chi.URLParam(r, "value"),
	}

	err := h.services.MetricsService.Add(m)
	if err != nil {
		if errors.Is(err, services.ErrIncorrectType) {
			w.WriteHeader(http.StatusNotImplemented)
			fmt.Fprint(w, "method not implemented")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "unknown error")
		return
	}

}
