package main

import (
	"time"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/cashe"
	"github.com/rusalexch/metal/internal/metric"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

func main() {
	// var url string
	// var port int

	m := metric.New()
	c := cashe.New()

	conf := app.Config{
		PollInterval:   pollInterval,
		ReportInterval: reportInterval,
		// ServerUrl:      url,
		// ServerPort:     port,
		Metrics: m,
		Cache:   c,
	}

	a := app.New(conf)

	a.Start()
}
