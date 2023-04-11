package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// New конструктор хранилища метрик
func New(dbURL *string) *Storage {
	st := &Storage{
		counters: map[string]MetricCounter{},
		gauges:   map[string]MetricGauge{},
	}
	if dbURL != nil {
		db, err := pgxpool.New(context.Background(), *dbURL)
		if err != nil {
			log.Panic(err)
		}
		st.db = db
	}
	return st
}

// AddCounter метод добавления метрики типа counter
func (s *Storage) AddCounter(name string, value int64) error {
	exist, _ := s.GetCounter(name)
	m := MetricCounter{
		Value: exist + value,
		Name:  name,
	}
	s.counters[name] = m

	return nil
}

// AddGuage метода добавления метрики типа guage
func (s *Storage) AddGauge(name string, value float64) error {
	m := MetricGauge{
		Value: value,
		Name:  name,
	}
	s.gauges[name] = m

	return nil
}

// GetCounter метод получения метрики типа counter
func (s *Storage) GetCounter(name string) (int64, error) {
	if m, ok := s.counters[name]; ok {
		return m.Value, nil
	}

	return 0, ErrMetricNotFound
}

// GetGuage метод получения метрики типа guage
func (s *Storage) GetGauge(name string) (float64, error) {
	if m, ok := s.gauges[name]; ok {
		return m.Value, nil
	}

	return 0, ErrMetricNotFound
}

// ListCounter метод получения списка метрик типа counter
func (s *Storage) ListCounter() []MetricCounter {
	if s.counters == nil {
		return []MetricCounter{}
	}
	res := make([]MetricCounter, 0, len(s.counters))

	for _, val := range s.counters {
		res = append(res, val)
	}

	return res
}

// ListCounter метод получения списка метрик типа gauge
func (s *Storage) ListGauge() []MetricGauge {
	if s.gauges == nil {
		return []MetricGauge{}
	}
	res := make([]MetricGauge, 0, len(s.gauges))

	for _, val := range s.gauges {
		res = append(res, val)
	}

	return res
}

func (s *Storage) Ping() error {
	if s.db != nil {
		return s.db.Ping(context.Background())
	}

	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}
