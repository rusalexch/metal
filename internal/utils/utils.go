package utils

import "strconv"

// isSameEmpty проверка слайса на пустые значение
// если хоть один элемент пустая строка то вернет true.
func IsSameEmpty(s []string) bool {
	res := false
	for _, item := range s {
		if item == "" {
			res = true
		}
	}

	return res
}

// StrToFloat64 - преобразование строки в float64.
func StrToFloat64(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

// StrToInt64 - преобразование строки в int64.
func StrToInt64(v string) (int64, error) {
	return strconv.ParseInt(v, 10, 64)
}

// Float64ToStr - преобразование float64 в строку.
func Float64ToStr(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

// Int64ToStr - преобразование int64 в строку.
func Int64ToStr(v int64) string {
	return strconv.FormatInt(v, 10)
}
