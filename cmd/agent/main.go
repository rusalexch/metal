package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rusalexch/metal/internal/agent"
	"github.com/rusalexch/metal/internal/cache"
	"github.com/rusalexch/metal/internal/config"
	"github.com/rusalexch/metal/internal/hash"
	"github.com/rusalexch/metal/internal/poll"
	"github.com/rusalexch/metal/internal/transport"
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
	env := config.NewAgentConfig()

	p := poll.New()
	c := cache.New()
	t := transport.New(env.Addr, env.RateLimit, env.PublicKey, env.GRPCAddress)
	h := hash.New(env.HashKey)

	conf := agent.Config{
		PollInterval:   env.PoolInterval,
		ReportInterval: env.ReportInterval,
		Poll:           p,
		Cache:          c,
		Transport:      t,
		Hash:           h,
	}

	a := agent.New(conf)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	a.Start(ctx)

	<-ctx.Done()
	log.Println("graceful shutdown")
}

func buildLog() {
	fmt.Printf("Build version: %s\n", utils.StringTernar(buildVersion, notValue))
	fmt.Printf("Build version: %s\n", utils.StringTernar(buildDate, notValue))
	fmt.Printf("Build version: %s\n", utils.StringTernar(buildCommit, notValue))
}
