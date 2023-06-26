package hash

import "hash"

// Hash - структура хэш-функции
type Hash struct {
	hash.Hash
	// needHash - флаг задействования хэш-функции
	needHash bool
}
