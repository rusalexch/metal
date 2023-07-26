package transport

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/utils"
)

// New - конструктор клиента отправки метрик.
func New(addr string, rateLimit int, publicKey *rsa.PublicKey) *Client {
	client := &http.Client{}

	return &Client{
		addr:      addr,
		client:    client,
		chOne:     make(chan app.Metrics),
		chJSONOne: make(chan app.Metrics),
		chList:    make(chan []app.Metrics),
		chReq:     make(chan reqParam),
		cntReq:    rateLimit,
		publicKey: publicKey,
	}
}

// Start - запуск клиента отправки метрик.
func (c *Client) Start(ctx context.Context, ch <-chan []app.Metrics) {
	c.init()

	for {
		select {
		case <-ctx.Done():
			c.close()
			return
		case m := <-ch:
			c.dmx(m)
		}
	}
}

// close - закрытие клиента отправки метрик.
func (c *Client) close() {
	close(c.chOne)
	close(c.chJSONOne)
	close(c.chList)
	close(c.chReq)
}

// dmx - демультиплексор метрик по каналам.
func (c *Client) dmx(m []app.Metrics) {
	go func() {
		for _, v := range m {
			if c.chOne != nil {
				c.chOne <- v
			}
		}
	}()
	go func() {
		for _, v := range m {
			if c.chJSONOne != nil {
				c.chJSONOne <- v
			}
		}
	}()
	go func() {
		if c.chList != nil {
			c.chList <- m
		}
	}()
}

// init - инициализация клиента отправки метрик.
func (c *Client) init() {
	c.initSendOne()
	c.initSendJSONOne()
	c.initSendList()
	for i := 0; i < c.cntReq; i++ {
		go func() {
			for r := range c.chReq {
				c.makeRequest(r)
			}
		}()
	}
}

// makeRequest - конструктор запроса отправки метрик на сервер.
func (c *Client) makeRequest(param reqParam) {
	body := bytes.NewBuffer(c.encrypt(param.body))
	req, err := http.NewRequest(http.MethodPost, param.url, body)
	if err != nil {
		log.Println(err)
	}
	if param.body == nil {
		req.Header.Add("Content-Type", "text/plain")
	} else {
		req.Header.Add("Content-Type", "application/json")

	}
	res, err := c.client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
}

// encrypt шифрование данных публичным ключом
func (c *Client) encrypt(body []byte) []byte {
	if c.publicKey == nil || body == nil {
		return body
	}

	encryptBody, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, c.publicKey, body, nil)
	if err != nil {
		log.Println("can't encrypt body")
		return body
	}

	return encryptBody
}

// initSendOne - инициализация канала отправки метрик по одиночному каналу.
func (c *Client) initSendOne() {
	if c.chOne == nil {
		return
	}
	go func() {
		for m := range c.chOne {
			c.chReq <- c.sendOneParam(m)
		}
	}()
}

// sendOneParam - подготовка параметров запроса одиной метрики.
func (c *Client) sendOneParam(m app.Metrics) reqParam {
	var val string
	switch m.Type {
	case app.Counter:
		val = utils.Int64ToStr(*m.Delta)
	case app.Gauge:
		val = utils.Float64ToStr(*m.Value)
	}
	url := fmt.Sprintf("http://%s/update/%s/%s/%s", c.addr, m.Type, m.ID, val)

	return reqParam{url: url, body: nil}
}

// initSendJSONOne - инициализация канала отправки одной метрики форматом JSON.
func (c *Client) initSendJSONOne() {
	if c.chJSONOne == nil {
		return
	}
	go func() {
		for m := range c.chJSONOne {
			c.chReq <- c.sendJSONOneParam(m)
		}
	}()
}

// sendJSONOneParam - подготовка параметров запроса отправки одной метрики форматом JSON.
func (c *Client) sendJSONOneParam(m app.Metrics) reqParam {
	url := fmt.Sprintf("http://%s/update/", c.addr)
	body, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}

	return reqParam{url: url, body: body}
}

// initSendList - инициализация канала отправки метрик списком.
func (c *Client) initSendList() {
	if c.chList == nil {
		return
	}
	go func() {
		for m := range c.chList {
			c.chReq <- c.sendListParam(m)
		}
	}()
}

// sendListParam - подготовка параметров отправки метрик списком.
func (c *Client) sendListParam(m []app.Metrics) reqParam {
	url := fmt.Sprintf("http://%s/updates/", c.addr)
	body, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}

	return reqParam{url: url, body: body}
}
