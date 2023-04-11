package services

import (
	"github.com/rusalexch/metal/internal/storage"
)

// Services структура сервисов
type Services struct {
	MetricsService Mertrics
	HealthCheck    Healther
}

// MertricsService структура сервиса метрик
type MertricsService struct {
	storage     storage.MetricsStorage
	subscribers []func()
}

type HealthCheck struct {
	repo Repositorier
}
