package cashe

import "github.com/rusalexch/metal/internal/agent"

// New инициализация кэша
func New() *Cashe {
	return &Cashe{}
}

// Add добавление значений метрик в кэш
func (c *Cashe) Add(m []agent.Metric) error {
	if c.m == nil {
		c.m = make([]agent.Metric, len(m))
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
func (c Cashe) Get() ([]agent.Metric, error) {
	return c.m, nil
}
