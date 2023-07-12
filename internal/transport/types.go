package transport

import (
	"io"
	"net/http"

	"github.com/rusalexch/metal/internal/app"
)

// Client - структура клиента.
type Client struct {
	// client http клиент.
	client *http.Client
	// chOne - канал отправки одной метрики.
	chOne chan app.Metrics
	// chJSONOne - канал отправки одной метрики формата JSON.
	chJSONOne chan app.Metrics
	// chList - канал отправки списка метрик.
	chList chan []app.Metrics
	// chReq - канал параметров запроса отправки метрик.
	chReq chan reqParam
	// addr адрес сервера сбора метрик.
	addr string
	// cntReq - количество одновременно запущенных сессий отправки метрик.
	cntReq int
}

// reqParam - структура параметров запроса на отправку метрик.
type reqParam struct {
	// body - данные для отправки, если есть.
	body io.Reader
	// url - url адрес отправки метрик.
	url string
}
