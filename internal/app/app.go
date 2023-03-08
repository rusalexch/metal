package app

import (
	"fmt"
	"log"
	"time"
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
func (a *App) Start() {
	pollTicker := time.NewTicker(a.pollInterval)
	defer pollTicker.Stop()
	reportTicker := time.NewTicker(a.reportInterval)
	defer reportTicker.Stop()

	for {
		select {
		case <-pollTicker.C:
			{
				a.scanAndSave()
				fmt.Println("poll")
			}
		case <-reportTicker.C:
			{
				a.send()
				fmt.Println("report")
			}
		}
	}
}

func (a *App) scanAndSave() {
	m, err := a.metrics.Scan()
	if err != nil {
		log.Fatal(err) //TODO
	}
	a.save(m)
}

func (a *App) save(m []Metric) {
	a.cache.Add(m)
}

func (a *App) send() {
	if a.transport == nil {
		log.Fatal("transport not provided")
	}
	list, err := a.cache.Get()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range list {
		err := a.transport.SendOne(item)
		if err != nil {
			log.Println(err)
		}
	}
}
