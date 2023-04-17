package app

import "errors"

var ErrNotFound = errors.New("metric not found")
var ErrIncorrectType = errors.New("incorrect metric type")
