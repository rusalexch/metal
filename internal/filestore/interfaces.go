package filestore

import "github.com/rusalexch/metal/internal/app"

type Saver interface {
	Download([]app.Metrics) error
	Upload() ([]app.Metrics, error)
}
