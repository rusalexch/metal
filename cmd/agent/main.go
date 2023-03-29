package main

import (
	"time"

	"github.com/rusalexch/metal/internal/agent"
	"github.com/rusalexch/metal/internal/cache"
	"github.com/rusalexch/metal/internal/config"
	"github.com/rusalexch/metal/internal/metric"
	"github.com/rusalexch/metal/internal/transport"
)


func main() {
	env := config.NewAgentConfig()

	m := metric.New()
	c := cache.New()
	t := transport.New(env.Addr)

	conf := agent.Config{
		PollInterval:   time.Second * time.Duration(env.PoolInterval),
		ReportInterval: time.Second * time.Duration(env.ReportInterval),
		Metrics:        m,
		Cache:          c,
		Transport:      t,
	}

	a := agent.New(conf)

	a.Start()
}
