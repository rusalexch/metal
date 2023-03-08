package cashe

import "github.com/rusalexch/metal/internal/app"

// New инициализация кэша
func New() *Cashe {
	return &Cashe{}
}

// Add добавление значений метрик в кэш
func (c *Cashe) Add(m []app.Metric) error {
	if c.m == nil {
		c.m = make([]app.Metric, len(m))
		copy(c.m, m)

		return nil
	}
	c.m = append(c.m, m...)
	return nil
}

// Reset сброс кэша
func (c *Cashe) Reset() error {
	c.m = nil
	return nil
}

// Get получение текущих значений кэша
func (c Cashe) Get() ([]app.Metric, error) {
	return c.m, nil
}
