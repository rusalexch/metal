package storage

import "errors"

var (
	// CounterNotFoundErr метрика типа counter не найдена
	ErrCounterNotFound = errors.New("counter metric not found")
	// CounterNotFoundErr метрика типа guage не найдена
	ErrGiuageNotFound = errors.New("guage metric not found")
)
