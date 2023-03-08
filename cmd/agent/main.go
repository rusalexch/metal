package main

import (
	"time"

	"github.com/rusalexch/metal/internal/app"
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

	conf := app.Config{
		PollInterval:   pollInterval,
		ReportInterval: reportInterval,
		// ServerUrl:      url,
		// ServerPort:     port,
		Metrics: m,
	}

	a := app.New(conf)

	a.Start()
}
