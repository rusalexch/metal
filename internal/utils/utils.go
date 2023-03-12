package utils

// isSameEmpty проверка слайса на пустые значение
// если хоть один элемент пустая строка то вернет true
func IsSameEmpty(s []string) bool {
	res := false
	for _, item := range s {
		if item == "" {
			res = true
		}
	}

	return res
}
