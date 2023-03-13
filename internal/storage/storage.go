package storage

// New конструктор хранилища метрик
func New() *Storage {
	return &Storage{
		counters: map[string]MetricCounter{},
		guages:   map[string]MetricGuage{},
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
func (s *Storage) AddGuage(name string, value float64) error {
	m := MetricGuage{
		Value: value,
		Name:  name,
	}
	s.guages[name] = m

	return nil
}

// GetCounter метод получения метрики типа counter
func (s *Storage) GetCounter(name string) (int64, error) {
	if m, ok := s.counters[name]; ok {
		return m.Value, nil
	}

	return 0, ErrCounterNotFound
}

// GetGuage метод получения метрики типа guage
func (s *Storage) GetGuage(name string) (float64, error) {
	if m, ok := s.guages[name]; ok {
		return m.Value, nil
	}

	return 0, ErrGiuageNotFound
}
