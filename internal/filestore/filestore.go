package filestore

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"time"

	"github.com/rusalexch/metal/internal/services"
)

func New(file string, interval time.Duration, restore bool, metricsService services.Mertrics) *Store {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(f)
	decoder := json.NewDecoder(f)

	return &Store{
		file:           f,
		encoder:        encoder,
		decoder:        decoder,
		storeInterval:  interval,
		restore:        restore,
		metricsService: metricsService,
	}
}

func (s *Store) Start() {
	if s.restore {
		s.upload()
	}
	if s.storeInterval == 0 {
		s.metricsService.Subscribe(s.download)
	} else {
		go s.start()
	}
}

func (s *Store) Close() error {
	return s.file.Close()
}

func (s *Store) download() {
	m := s.metricsService.List()
	store := StoreMetrics{
		Metrics: m,
	}
	err := s.encoder.Encode(store)
	if err != nil {
		log.Println(err)
	}
}

func (s *Store) upload() {
	var metrics StoreMetrics
	err := s.decoder.Decode(&metrics)
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatal(err)
	}
	for _, m := range metrics.Metrics {
		s.metricsService.Add(m)
	}
}

func (s *Store) start() {
	ticker := time.NewTicker(s.storeInterval)
	defer ticker.Stop()
	for {
		<-ticker.C
		s.download()
	}
}
