package main

import (
	"fmt"

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
	t := transport.New(env.Addr, env.RateLimit)
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

	a.Start()
}

func buildLog() {
	fmt.Printf("Build version: %s\n", utils.StringTernar(buildVersion, notValue))
	fmt.Printf("Build version: %s\n", utils.StringTernar(buildDate, notValue))
	fmt.Printf("Build version: %s\n", utils.StringTernar(buildCommit, notValue))
}
