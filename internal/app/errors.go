package app

import "errors"

var (
	// ErrNotFound - ошибка: метрика не найдена
	ErrNotFound = errors.New("metric not found")
	// ErrIncorrectType - ошибка: не корректный тип метрики
	ErrIncorrectType = errors.New("incorrect metric type")
)
