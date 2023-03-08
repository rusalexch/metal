package main

import (
	"time"

	"github.com/rusalexch/metal/internal/agent"
	"github.com/rusalexch/metal/internal/cashe"
	"github.com/rusalexch/metal/internal/metric"
	"github.com/rusalexch/metal/internal/transport"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

func main() {
	var url string
	var port int

	m := metric.New()
	c := cashe.New()
	t := transport.New(url, port)

	conf := agent.Config{
		PollInterval:   pollInterval,
		ReportInterval: reportInterval,
		Metrics:        m,
		Cache:          c,
		Transport:      t,
	}

	a := agent.New(conf)

	a.Start()
}
