package cache

import (
	"context"
	"time"

	"github.com/rusalexch/metal/internal/app"
)

// New - конструктор кэша.
func New() *Cache {
	return &Cache{
		m: map[string]app.Metrics{},
	}
}

// Start - запуск кеша.
func (c *Cache) Start(ctx context.Context, chIn <-chan app.Metrics, chOut chan<- []app.Metrics, t time.Ticker) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case m := <-chIn:
				c.add(m)
			case <-t.C:
				chOut <- c.get()
			}
		}
	}()
}

// add - метод добавления метрики в кэш.
func (c *Cache) add(m app.Metrics) {
	c.Lock()
	defer c.Unlock()

	c.m[m.ID] = m
}

// get - метод получения метрики.
func (c *Cache) get() []app.Metrics {
	c.Lock()
	defer c.Unlock()

	m := make([]app.Metrics, 0, len(c.m))
	for _, v := range c.m {
		m = append(m, v)
	}
	c.m = map[string]app.Metrics{}

	return m
}
