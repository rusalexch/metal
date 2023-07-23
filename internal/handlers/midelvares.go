package handlers

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// gzipWriter - структура миддлвара сжатия.
type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

// Write - реализация интерфейса io.Writer.
func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// compressMiddleware - middleware сжатия ответа.
func compressMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

// decompressMiddleware - middleware распаковки сжатого запроса.
func decompressMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Encoding") == "gzip" {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer gz.Close()
			r.Body = gz
		}
		defer r.Body.Close()

		next.ServeHTTP(w, r)
	})
}

func (h *Handlers) decryptMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h.privateKey == nil {
			next.ServeHTTP(w, r)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("decryptMiddleware > can't read body")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		decryptBody, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, h.privateKey, body, nil)
		if err != nil {
			log.Println("decryptMiddleware > can't decrypt body")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewReader(decryptBody))
		next.ServeHTTP(w, r)
	})
}
