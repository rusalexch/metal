package config

import "time"

const (
	defaultAddr           = "127.0.0.1:8080"
	defaultReportInterval = time.Second * 10
	defaultPoolInterval   = time.Second * 2
	defaultRestore        = "true"
	defaultStoreInterval  = time.Second * 300
	defaultStoreFile      = "/tmp/devops-metrics-db.json"
	defaultKey            = ""
	defaultRateLimit      = 1
)
