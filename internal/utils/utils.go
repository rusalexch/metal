package utils

import "strconv"

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

func StrToFloat64(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

func StrToInt64(v string) (int64, error) {
	return strconv.ParseInt(v, 10, 64)
}

func Float64ToStr(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func Int64ToStr(v int64) string {
	return strconv.FormatInt(v, 10)
}
