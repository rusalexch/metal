package storage

// New конструктор хранилища метрик
func New() *Storage {
	return &Storage{
		counters: map[string]MetricCounter{},
		gauges:   map[string]MetricGauge{},
	}
}

// AddCounter метод добавления метрики типа counter
func (s *Storage) AddCounter(name string, value int64) error {
	m := MetricCounter{
		Value: value,
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
