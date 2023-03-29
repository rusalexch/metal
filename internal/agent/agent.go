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
func New(conf Config) *Agent {

	return &Agent{
		pollInterval:   conf.PollInterval,
		reportInterval: conf.ReportInterval,
		metrics:        conf.Metrics,
		cache:          conf.Cache,
		transport:      conf.Transport,
	}
}

// Start метод запуска клиента сбора и отправки метрик на сервер
func (a *Agent) Start() error {
	pollTicker := time.NewTicker(a.pollInterval)
	defer pollTicker.Stop()
	reportTicker := time.NewTicker(a.reportInterval)
	defer reportTicker.Stop()

	for {
		select {
		case <-pollTicker.C:
			{
				a.scanAndSave()
			}
		case <-reportTicker.C:
			{
				err := a.send()
				if err != nil {
					log.Println(err)
				}
				a.cache.Reset()
			}
		}
	}
}

func (a *Agent) scanAndSave() {
	m := a.metrics.Scan()

	a.save(m)
}

func (a *Agent) save(m []app.Metrics) {
	a.cache.Add(m)
}

func (a *Agent) send() error {
	if a.transport == nil {
		return errors.New(TransportNotProvided)
	}
	list := a.cache.Get()

	isError := false
	for _, item := range list {
		err := a.transport.SendOne(item)
		if err != nil {
			log.Println(err)
			isError = true
		} else {
			log.Println(fmt.Scanf("metric: %s was sended", item.ID))
		}
		err = a.transport.SendOneJSON(item)
		if err != nil {
			log.Println(err)
			isError = true
		} else {
			log.Println(fmt.Scanf("metric: %s was sended", item.ID))
		}
	}
	if isError {
		return errors.New(NotAllMetricsSent)
	}
	return nil
}