package main

import (
	"github.com/rusalexch/metal/internal/agent"
	"github.com/rusalexch/metal/internal/cache"
	"github.com/rusalexch/metal/internal/config"
	"github.com/rusalexch/metal/internal/hash"
	"github.com/rusalexch/metal/internal/metric"
	"github.com/rusalexch/metal/internal/transport"
)

func main() {
	env := config.NewAgentConfig()

	m := metric.New()
	c := cache.New()
	t := transport.New(env.Addr)
	h := hash.New(env.HashKey)

	conf := agent.Config{
		PollInterval:   env.PoolInterval,
		ReportInterval: env.ReportInterval,
		Metrics:        m,
		Cache:          c,
		Transport:      t,
		Hash:        h,
	}

	a := agent.New(conf)

	a.Start()
}
