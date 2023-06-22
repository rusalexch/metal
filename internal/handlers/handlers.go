package handlers

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
)

// New конструктор Хэндлераов
func New(stor storager, h hasher) *Handlers {
	return &Handlers{
		storage: stor,
		hash:    h,
		timeout: 100 * time.Second,
		Mux:     chi.NewMux(),
	}
}

// Init инициализация Хендлеров
func (h *Handlers) Init() {
	logger := httplog.NewLogger("httplog", httplog.Options{
		JSON: true,
	})

	h.Use(middleware.RequestID)
	h.Use(middleware.RealIP)
	h.Use(httplog.RequestLogger(logger))
	h.Use(compressMiddleware)
	h.Use(decompressMiddleware)
	h.Use(middleware.Recoverer)

	h.Get("/", h.list)
	h.Get("/ping", h.ping)
	h.Get("/value/{mType}/{ID}", h.find)
	h.Post("/update/{mType}/{ID}/{value}", h.update)
	h.Post("/update/", h.updateJSON)
	h.Post("/value/", h.valueJSON)
	h.Post("/updates/", h.updates)
	h.Mount("/debug", middleware.Profiler())
}
