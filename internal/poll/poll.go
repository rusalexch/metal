package poll

import (
	"time"

	"github.com/rusalexch/metal/internal/metric"
)

type Metrics interface {
	Get() metric.Metrics
}

type Poll struct {
	PollInterval time.Duration
	m            Metrics
}

func New(pollInterval time.Duration, m Metrics) *Poll {
	return &Poll{
		PollInterval: pollInterval,
		m:            m,
	}
}
