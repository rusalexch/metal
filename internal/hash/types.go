package hash

import "hash"

type Hash struct {
	hash.Hash
	// needHash статус разрешенного хеширования
	needHash bool
}
