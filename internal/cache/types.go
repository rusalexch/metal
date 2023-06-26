package cache

import (
	"sync"

	"github.com/rusalexch/metal/internal/app"
)

// Cache - структура кэша
type Cache struct {
	// слайс для хранения кэша
	m map[string]app.Metrics
	*sync.Mutex
}
