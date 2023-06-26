package utils

// Int64AsPointer - число int64 в указатель
func Int64AsPointer(v int64) *int64 {
	return &v
}

// Float64AsPointer - число float64 в указатель
func Float64AsPointer(v float64) *float64 {
	return &v
}
