package hash

import "hash"

type Hash struct {
	hash.Hash
	// isEnable статус разрешенного хеширования
	isEnable bool
}
