package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/utils"
)

// update Хэндлер для обновления метрик
func (h *Handlers) update(w http.ResponseWriter, r *http.Request) {
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

	err := h.storage.Add(m)
	if err != nil {
		if errors.Is(err, app.ErrIncorrectType) {
			w.WriteHeader(http.StatusNotImplemented)
			fmt.Fprint(w, "method not implemented")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "unknown error")
		return
	}

}

func (h *Handlers) updateJSON(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("readAll body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	content := r.Header.Get(contentType)
	if content != appJSON {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	var m app.Metrics
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Println("unmarshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isCheck := h.hash.Check(m); !isCheck {
		log.Println("invalid hash")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.storage.Add(m)

	m, _ = h.storage.Get(m.ID, m.Type)
	h.hash.AddHash(&m)
	body, err = json.Marshal(m)
	if err != nil {
		log.Println("marshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add(contentType, appJSON)
	w.Write(body)
}

func (h *Handlers) updates(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("readAll body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	content := r.Header.Get(contentType)
	if content != appJSON {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	var m []app.Metrics
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Println("unmarshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(m)

	err = h.storage.AddList(m)
	if err != nil {
		log.Println("addList", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	m, err = h.storage.List()
	if err != nil {
		log.Println("getList", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err = json.Marshal(m)
	if err != nil {
		log.Println("marshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add(contentType, appJSON)
	w.Write(body)
}
