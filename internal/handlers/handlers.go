package handlers

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/rusalexch/metal/internal/services"
)

// New конструктор Хэндлераов
func New(services *services.Services) *Handlers {
	return &Handlers{
		services: services,
		Mux:      chi.NewMux(),
	}
}

// Init инициализация Хендлеров
func (h *Handlers) Init() {
	h.Use(middleware.RequestID)
	h.Use(middleware.RealIP)
	h.Use(middleware.Logger)
	h.Use(middleware.Recoverer)
	h.Use(compressMiddleware)
	h.Use(decompressMiddleware)

	h.Get("/", h.list)
	h.Get("/ping", ping)
	h.Get("/value/{mType}/{ID}", h.find)
	h.Post("/update/{mType}/{ID}/{value}", h.update)
	h.Post("/update/", h.updateJSON)
	h.Post("/value/", h.valueJSON)
}
