package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/rusalexch/metal/internal/app"
)

func New(key string) *Hash {
	return &Hash{
		Hash:     hmac.New(sha256.New, []byte(key)),
		needHash: key == "",
	}
}

func (h Hash) AddHash(m *app.Metrics) {
	if h.needHash {
		return
	}

	m.Hash = h.createHash(m)
}

func (h Hash) Check(m app.Metrics) bool {
	if h.needHash {
		return true
	}
	checkHash := h.createHash(&m)

	return checkHash == m.Hash
}

func (h Hash) createHash(m *app.Metrics) string {
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
	log.Printf("%x\n", hash)
	return fmt.Sprintf("%x", hash)
}
