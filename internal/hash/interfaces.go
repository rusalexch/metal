package hash

import "github.com/rusalexch/metal/internal/app"

type Hasher interface {
	AddHash(m *app.Metrics)
	Check(m app.Metrics) bool
}
