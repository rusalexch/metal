package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/services"
	"github.com/rusalexch/metal/internal/utils"
)

// update Хэндлер для обновления метрик
func (h *Handlers) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "method not available")
		return
	}

	// if r.Header.Get("Content-Type") != "text/plain" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprint(w, "Content-Type not available")
	// 	return
	// }

	var data string
	fmt.Sscanf(r.URL.String(), "/update/%s/%s/%s", &data)
	s := strings.Split(data, "/")
	if len(s) != 3 || utils.IsSameEmpty(s) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "required three params")
		return
	}

	m := app.Metric{
		Type:      s[0],
		Name:      s[1],
		Value:     s[2],
		Timestamp: 0,
	}

	err := h.services.MetricsService.Add(m)
	if err != nil {
		if errors.Is(err, services.IncorrectTypeErr) {
			w.WriteHeader(http.StatusNotImplemented)
			fmt.Fprint(w, "method mot implemented")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "unknown error")
		return
	}

}
