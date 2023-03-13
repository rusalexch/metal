package handlers

import (
	"net/http"

	"github.com/rusalexch/metal/internal/services"
)

// New конструктор Хэндлераов
func New(services *services.Services) *Handlers {
	return &Handlers{
		services: services,
	}
}

// Init инициализация Хендлеров
func (h *Handlers) Init() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/update/", h.update)
}
