package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/rusalexch/metal/internal/app"
)

func New(key string) *Hash {
	isEnable := key != ""
	return &Hash{
		hmac.New(sha256.New, []byte(key)),
		isEnable,
	}
}

func (h Hash) AddHash(m *app.Metrics) {
	if !h.isEnable {
		return
	}

	m.Hash = h.createHash(*m)
}

func (h Hash) Check(m app.Metrics) bool {
	if !h.isEnable {
		return true
	}
	checkHash := h.createHash(m)

	return checkHash == m.Hash
}

func (h Hash) createHash(m app.Metrics) string {
	str := ""
	switch m.Type {
	case app.Counter:
		str = fmt.Sprintf("%s:counter:%d", m.ID, *m.Delta)
	case app.Gauge:
		str = fmt.Sprintf("%s:gauge:%f", m.ID, *m.Value)
	}
	_, err := h.Write([]byte(str))
	if err != nil {
			log.Println("addHash error:", err)
	}
	hash := h.Sum(nil)
	return string(hash)
}
