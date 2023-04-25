package cache

import (
	"context"
	"time"

	"github.com/rusalexch/metal/internal/app"
)

// New инициализация кэша
func New() *Cache {
	return &Cache{
		m: map[string]app.Metrics{},
	}
}

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

func (c *Cache) add(m app.Metrics) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.m[m.ID] = m
}

func (c *Cache) get() []app.Metrics {
	c.mx.Lock()
	defer c.mx.Unlock()

	m := make([]app.Metrics, 0, len(c.m))
	for _, v := range c.m {
		m = append(m, v)
	}
	c.m = map[string]app.Metrics{}

	return m
}
