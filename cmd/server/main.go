package main

import (
	"log"

	"github.com/rusalexch/metal/internal/config"
	"github.com/rusalexch/metal/internal/filestore"
	"github.com/rusalexch/metal/internal/handlers"
	"github.com/rusalexch/metal/internal/hash"
	"github.com/rusalexch/metal/internal/server"
	"github.com/rusalexch/metal/internal/services"
	"github.com/rusalexch/metal/internal/storage"
)

func main() {
	envConf := config.NewServerConfig()
	stor := storage.New(&envConf.DBURL)
	defer stor.Close()
	srv := services.New(stor)
	hs := hash.New(envConf.HashKey)
	h := handlers.New(srv, hs)
	s := server.New(h, envConf.Addr)

	fs := filestore.New(envConf.StoreFile, envConf.StoreInterval, envConf.Restore, srv.MetricsService)
	defer fs.Close()
	fs.Start()

	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
