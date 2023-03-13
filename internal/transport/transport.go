package transport

import (
	"fmt"
	"net/http"

	"github.com/rusalexch/metal/internal/app"
)

// New конструктор клиента
func New(addr string, port int) *Client {
	if addr == "" {
		addr = defaultAddr
	}
	if port == 0 {
		port = defaultPort
	}

	client := &http.Client{}

	return &Client{
		addr:   addr,
		port:   port,
		client: client,
	}
}

// SendOne отправка одной метрики на сервер
func (c *Client) SendOne(m app.Metric) error {
	req, err := http.NewRequest(http.MethodPost, c.url(m), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

// url метод получения url
func (c Client) url(m app.Metric) string {
	return fmt.Sprintf("%s:%d/update/%s/%s/%s", c.addr, c.port, m.Type, m.Name, m.Value)
}