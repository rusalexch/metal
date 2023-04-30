package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/utils"
)

// New конструктор клиента
func New(addr string, rateLimit int) *Client {
	client := &http.Client{}

	return &Client{
		addr:      addr,
		client:    client,
		chOne:     make(chan app.Metrics),
		chJSONOne: make(chan app.Metrics),
		chList:    make(chan []app.Metrics),
		chReq:     make(chan reqParam),
		cntReq:    rateLimit,
	}
}

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

func (c *Client) close() {
	close(c.chOne)
	close(c.chJSONOne)
	close(c.chList)
	close(c.chReq)
}

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

func (c *Client) makeRequest(param reqParam) {
	req, err := http.NewRequest(http.MethodPost, param.url, param.body)
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
	}
	defer res.Body.Close()
}

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

func (c *Client) sendJSONOneParam(m app.Metrics) reqParam {
	url := fmt.Sprintf("http://%s/update/", c.addr)
	body, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}

	return reqParam{url: url, body: bytes.NewBuffer(body)}
}

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

func (c *Client) sendListParam(m []app.Metrics) reqParam {
	url := fmt.Sprintf("http://%s/updates/", c.addr)
	body, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}

	return reqParam{url: url, body: bytes.NewBuffer(body)}
}
