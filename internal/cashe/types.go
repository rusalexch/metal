package cashe

import "github.com/rusalexch/metal/internal/agent"

// Cashe структура кэша
type Cashe struct {
	// слайс для хранения кэша
	m []agent.Metric
}
