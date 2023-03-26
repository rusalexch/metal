package cache

import "github.com/rusalexch/metal/internal/app"

// New инициализация кэша
func New() *Cache {
	return &Cache{}
}

// Add добавление значений метрик в кэш
func (c *Cache) Add(m []app.Metric) {
	if c.m == nil {
		c.m = make([]app.Metric, len(m))
		copy(c.m, m)
	}
	c.m = append(c.m, m...)
}

// Reset сброс кэша
func (c *Cache) Reset() {
	c.m = []app.Metric{}
}

// Get получение текущих значений кэша
func (c *Cache) Get() []app.Metric {
	if c.m == nil {
		c.m = []app.Metric{}
	}
	return c.m
}
