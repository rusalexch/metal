package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

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
	h := handlers.New(stor, hs, envConf.PrivateKey)
	s := server.New(h, envConf.Addr)

	// через этот канал сообщим основному потоку, что соединения закрыты
	idleConnsClosed := make(chan struct{})
	// канал для перенаправления прерываний
	// поскольку нужно отловить всего одно прерывание,
	// ёмкости 1 для канала будет достаточно
	sigint := make(chan os.Signal, 1)
	// регистрируем перенаправление прерываний
	signal.Notify(sigint, os.Interrupt)
	go func() {
		<-sigint
		s.Shutdown(context.Background(), idleConnsClosed)
	}()
	s.Start()

	<-idleConnsClosed
	log.Println("Server Shutdown gracefully")
}

func buildLog() {
	fmt.Printf("Build version: %s\n", utils.StringTernar(buildVersion, notValue))
	fmt.Printf("Build version: %s\n", utils.StringTernar(buildDate, notValue))
	fmt.Printf("Build version: %s\n", utils.StringTernar(buildCommit, notValue))
}
