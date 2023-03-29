package main

import (
	"log"

	"github.com/rusalexch/metal/internal/config"
	"github.com/rusalexch/metal/internal/handlers"
	"github.com/rusalexch/metal/internal/server"
	"github.com/rusalexch/metal/internal/services"
	"github.com/rusalexch/metal/internal/storage"
)

func main() {
	conf := config.NewServerConfig()
	stor := storage.New()
	srv := services.New(stor)
	h := handlers.New(srv)
	s := server.New(h, conf.Addr)

	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
