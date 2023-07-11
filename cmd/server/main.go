package main

import (
	"fmt"
	"log"

	"github.com/rusalexch/metal/internal/config"
	"github.com/rusalexch/metal/internal/handlers"
	"github.com/rusalexch/metal/internal/hash"
	"github.com/rusalexch/metal/internal/server"
	"github.com/rusalexch/metal/internal/storage"
	"github.com/rusalexch/metal/internal/utils"
)

const notValue = "N/A"

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

func main() {
	buildLog()
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

func buildLog() {
	fmt.Printf("Build version: %s\n", utils.StringTernar(buildVersion, notValue))
	fmt.Printf("Build version: %s\n", utils.StringTernar(buildDate, notValue))
	fmt.Printf("Build version: %s\n", utils.StringTernar(buildCommit, notValue))
}
