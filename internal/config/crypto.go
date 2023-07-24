package config

import (
	"crypto/rsa"
	"crypto/x509"
	"os"
)

// getPublicKey - получить публичный ключ
func getPublicKey(path *string) (*rsa.PublicKey, error) {
	if path == nil || *path == "" {
		return nil, nil
	}
	key, err := readCryptoFile(*path)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PublicKey(key)
}

// getPrivateKey - получить приватный ключ
func getPrivateKey(path *string) (*rsa.PrivateKey, error) {
	if path == nil || *path == "" {
		return nil, nil
	}
	key, err := readCryptoFile(*path)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(key)
}

// readCryptoFile - прочитать файл ключа
func readCryptoFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	return file, nil
}
