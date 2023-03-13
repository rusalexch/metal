package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/utils"
)

// update Хэндлер для обновления метрик
func (h *Handlers) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "method not available")
		return
	}

	if r.Header.Get("Content-Type") != "text/plain" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Content-Type not available")
		return
	}

	var data string
	fmt.Sscanf(r.URL.String(), "/update/%s/%s/%s", &data)
	s := strings.Split(data, "/")
	if len(s) != 3 || utils.IsSameEmpty(s) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "required three params")
		return
	}

	m := app.Metric{
		Type:      s[0],
		Value:     s[1],
		Name:      s[2],
		Timestamp: 0,
	}

	fmt.Println(m)

	h.services.MetricsService.Add(m)

}
