package transport

import "net/http"

// Client структура клиента
type Client struct {
	// addr адрес сервера сбора метрик
	addr string
	// client http клиент
	client *http.Client
}
