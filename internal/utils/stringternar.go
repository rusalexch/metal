package utils

// StringTernar - возвращает значение value если оно не пустое или значение по умолчанию def
func StringTernar(value, def string) string {
	if value == "" {
		return def
	}
	return value
}
