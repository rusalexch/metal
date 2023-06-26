package storage

import (
	"github.com/rusalexch/metal/internal/storage/internal/db"
	"github.com/rusalexch/metal/internal/storage/internal/filestorage"
)

// New - конструктор хранилища метрик.
func New(dbURL string, file string, restore bool) *Storage {
	var stor metricsStorage
	if dbURL != "" {
		stor = db.New(dbURL)
	} else {
		stor = filestorage.New(file, restore)
	}

	return &Storage{
		stor,
	}
}
