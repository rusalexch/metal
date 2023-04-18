package filestorage

import (
	"github.com/rusalexch/metal/internal/app"
)

type store struct {
	Counters map[string]int64   `json:"counters"`
	Gauges   map[string]float64 `json:"gauges"`
}

func (st *store) addMetric(m app.Metrics) {
	if m.Type == app.Counter {
		delta, isExist := st.Counters[m.ID]
		if isExist {
			st.Counters[m.ID] = delta + *m.Delta
		} else {
			st.Counters[m.ID] = *m.Delta
		}
	}
	if m.Type == app.Gauge {
		st.Gauges[m.ID] = *m.Value
	}
}
