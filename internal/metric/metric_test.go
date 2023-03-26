package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Metrics
	}{
		{
			name: "should be created",
			want: &Metrics{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMetrics_Scan(t *testing.T) {
	mNames := []string{
		"Alloc",
		"BuckHashSys",
		"Frees",
		"GCCPUFraction",
		"GCSys",
		"HeapAlloc",
		"HeapIdle",
		"HeapInuse",
		"HeapObjects",
		"HeapReleased",
		"HeapSys",
		"LastGC",
		"Lookups",
		"MCacheInuse",
		"MCacheSys",
		"MSpanInuse",
		"MSpanSys",
		"Mallocs",
		"NextGC",
		"NumForcedGC",
		"NumGC",
		"OtherSys",
		"PauseTotalNs",
		"StackInuse",
		"StackInuse",
		"StackSys",
		"Sys",
		"TotalAlloc",
		"PollCount",
		"RandomValue",
	}
	m := New()
	got := m.Scan()

	assert.Equal(t, 30, len(got))
	for _, item := range got {
		assert.Contains(t, mNames, item.Name)
		if item.Name == "PollCount" {
			assert.Equal(t, item.Value, "1")
		}
	}
}
