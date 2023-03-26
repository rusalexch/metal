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

	h.Get("/", h.list)
	h.Get("/ping", ping)
	h.Get("/value/{mType}/{name}", h.find)
	h.Post("/update/{mType}/{name}/{value}", h.update)
}
