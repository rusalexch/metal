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
				return a.scanAndSave()
			}
		case <-reportTicker.C:
			{
				fmt.Println("report")
				return a.send()
			}
		}
	}
}

func (a *App) scanAndSave() error {
	m := a.metrics.Scan()

	return a.save(m)
}

func (a *App) save(m []app.Metric) error {
	return a.cache.Add(m)
}

func (a *App) send() error {
	if a.transport == nil {
		return errors.New(TransportNotProvided)
	}
	list, err := a.cache.Get()
	if err != nil {
		return err
	}
	cntErr := 0
	for _, item := range list {
		err := a.transport.SendOne(item)
		if err != nil {
			log.Println(err)
			cntErr++
		} else {
			log.Println(fmt.Scanf("metric: %s was sended", item.Name))
		}
	}
	if cntErr == 0 {
		return nil
	}
	return errors.New(NotAllMetricsSent)
}
