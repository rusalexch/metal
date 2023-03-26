package cache

import "github.com/rusalexch/metal/internal/app"

// Cache структура кэша
type Cache struct {
	// слайс для хранения кэша
	m []app.Metric
}
