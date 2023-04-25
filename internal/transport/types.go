package transport

import (
	"io"
	"net/http"

	"github.com/rusalexch/metal/internal/app"
)

// Client структура клиента
type Client struct {
	// addr адрес сервера сбора метрик
	addr string
	// client http клиент
	client *http.Client

	chOne     chan app.Metrics
	chJsonOne chan app.Metrics
	chList    chan []app.Metrics
	chReq     chan reqParam
	cntReq    int
}

type reqParam struct {
	url  string
	body io.Reader
}
