package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/utils"
)

// New конструктор клиента
func New(addr string) *Client {
	client := &http.Client{}

	return &Client{
		addr:   addr,
		client: client,
	}
}

// SendOne отправка одной метрики на сервер
func (c *Client) SendOne(m app.Metrics) error {
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

func (c *Client) SendOneJSON(m app.Metrics) error {
	url := fmt.Sprintf("http://%s/update/", c.addr)
	body, err := json.Marshal(m)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

// url метод получения url
func (c Client) url(m app.Metrics) string {
	var val string
	switch m.Type {
	case app.Counter:
		val = utils.Int64ToStr(*m.Delta)
	case app.Gauge:
		val = utils.Float64ToStr(*m.Value)
	}
	return fmt.Sprintf("http://%s/update/%s/%s/%s", c.addr, m.Type, m.ID, val)
}
