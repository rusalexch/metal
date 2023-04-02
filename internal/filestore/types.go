package filestore

import (
	"encoding/json"
	"os"
	"time"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/services"
)

type Store struct {
	file           *os.File
	encoder        *json.Encoder
	decoder        *json.Decoder
	storeInterval  time.Duration
	restore        bool
	metricsService services.Mertrics
}

type StoreMetrics struct {
	Metrics []app.Metrics `json:"metrics"`
}
