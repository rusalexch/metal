package main

import (
	"log"

	"github.com/rusalexch/metal/internal/config"
	"github.com/rusalexch/metal/internal/handlers"
	"github.com/rusalexch/metal/internal/hash"
	"github.com/rusalexch/metal/internal/server"
	"github.com/rusalexch/metal/internal/storage"
)

func main() {
	envConf := config.NewServerConfig()
	stor := storage.New(envConf.DBURL, envConf.StoreFile, envConf.Restore)
	defer stor.Close()

	hs := hash.New(envConf.HashKey)
	h := handlers.New(stor, hs)
	s := server.New(h, envConf.Addr)

	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
