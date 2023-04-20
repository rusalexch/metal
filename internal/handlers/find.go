package handlers

import (
	"context"
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

func (h *Handlers) find(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	ID := chi.URLParam(r, "ID")
	mType := chi.URLParam(r, "mType")
	m, err := h.storage.Get(ctx, ID, mType)
	if err != nil {
		if errors.Is(err, app.ErrIncorrectType) {
			w.WriteHeader(http.StatusNotImplemented)
			fmt.Fprint(w, "method not implemented")
			return
		}
		if errors.Is(err, app.ErrNotFound) {
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

func (h *Handlers) valueJSON(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	var m app.Metrics

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

	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Println("unmarshal", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m, err = h.storage.Get(ctx, m.ID, m.Type)
	if err != nil {
		log.Println("get storage", err)
		if errors.Is(err, app.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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
