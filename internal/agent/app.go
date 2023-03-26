package agent

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/rusalexch/metal/internal/app"
)

// New инициализация приложения
// pollInterval - частота опроса метрик
// reportInterval - частота отправки метрик на сервер
// url - адрес сервера, по умолчанию "http://127.0.0.1"
// port - порт сервера, по умолчанию 8080
func New(conf Config) *App {

	return &App{
		pollInterval:   conf.PollInterval,
		reportInterval: conf.ReportInterval,
		metrics:        conf.Metrics,
		cache:          conf.Cache,
		transport:      conf.Transport,
	}
}

// Start метод запуска клиента сбора и отправки метрик на сервер
func (a *App) Start() error {
	pollTicker := time.NewTicker(a.pollInterval)
	defer pollTicker.Stop()
	reportTicker := time.NewTicker(a.reportInterval)
	defer reportTicker.Stop()

	for {
		select {
		case <-pollTicker.C:
			{
				fmt.Println("poll")
				a.scanAndSave()
			}
		case <-reportTicker.C:
			{
				fmt.Println("report")
				err := a.send()
				if err != nil {
					log.Println(err)
				}
				a.cache.Reset()
			}
		}
	}
}

func (a *App) scanAndSave() {
	m := a.metrics.Scan()

	a.save(m)
}

func (a *App) save(m []app.Metric) {
	a.cache.Add(m)
}

func (a *App) send() error {
	if a.transport == nil {
		return errors.New(TransportNotProvided)
	}
	list := a.cache.Get()

	cntErr := 0
	for _, item := range list {
		err := a.transport.SendOne(item)
		if err != nil {
			log.Println(err)
			cntErr++
		} else {
			log.Println(fmt.Scanf("metric: %s was sended", item.ID))
		}
	}
	if cntErr == 0 {
		return nil
	}
	return errors.New(NotAllMetricsSent)
}
