package cashe

import "github.com/rusalexch/metal/internal/app"

// Cashe структура кэша
type Cashe struct {
	// слайс для хранения кэша
	m []app.Metric
}
