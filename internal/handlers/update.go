package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/services"
	"github.com/rusalexch/metal/internal/utils"
)

// update Хэндлер для обновления метрик
func (h *Handlers) update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update")
	m := app.Metrics{
		Type: chi.URLParam(r, "mType"),
		ID:   chi.URLParam(r, "ID"),
	}
	switch m.Type {
	case app.Counter:
		{
			delta, err := utils.StrToInt64(chi.URLParam(r, "value"))
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "error counter value")
				return
			}
			m.Delta = &delta
		}
	case app.Gauge:
		{
			value, err := utils.StrToFloat64(chi.URLParam(r, "value"))
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "error gauge value")
				return
			}
			m.Value = &value
		}
	default:
		{
			w.WriteHeader(http.StatusNotImplemented)
			fmt.Fprint(w, "method not implemented")
			return
		}
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
