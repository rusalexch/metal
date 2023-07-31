package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rusalexch/metal/internal/hash"
	"github.com/rusalexch/metal/internal/storage"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (int, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	require.NoError(t, err)
	if body != nil {
		req.Header.Add(contentType, appJSON)
	}

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	defer resp.Body.Close()

	return resp.StatusCode, string(respBody)
}

func TestNew(t *testing.T) {
	h := New(storage.New("", "/tmp/test1", false), hash.New(""), nil, nil)
	h.Init()
	ts := httptest.NewServer(h)
	defer ts.Close()

	statusCode, body := testRequest(t, ts, http.MethodGet, "/ping", nil)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "", body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/", nil)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, htmlNoMetrics, body)

	statusCode, _ = testRequest(t, ts, http.MethodPost, "/update/gauge/testGauge1/0.000001", nil)
	assert.Equal(t, http.StatusOK, statusCode)

	statusCode, body = testRequest(t, ts, http.MethodPost, "/update/unknown/testGauge1/0.000001", nil)
	assert.Equal(t, http.StatusNotImplemented, statusCode)
	assert.Equal(t, "method not implemented", body)

	statusCode, _ = testRequest(t, ts, http.MethodPost, "/update/counter/testCounter1/100", nil)
	assert.Equal(t, http.StatusOK, statusCode)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/", nil)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, htmlWithMetrics, body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/value/gauge/testGauge1", nil)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "0.000001", body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/value/counter/testCounter1", nil)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "100", body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/value/gauge/testGauge11", nil)
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Equal(t, "metric not found", body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/value/counter/testCounter11", nil)
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Equal(t, "metric not found", body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/value/noname/testCounter11", nil)
	assert.Equal(t, http.StatusNotImplemented, statusCode)
	assert.Equal(t, "method not implemented", body)

	statusCode, body = testRequest(t, ts, http.MethodPost, "/update/", strings.NewReader(jsonCounter))
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, jsonCounter, body)

	statusCode, body = testRequest(t, ts, http.MethodPost, "/update/", strings.NewReader(jsonCounter))
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, jsonDoubleCounter, body)

	statusCode, body = testRequest(t, ts, http.MethodPost, "/update/", strings.NewReader(jsonGauge))
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, jsonGauge, body)

	statusCode, body = testRequest(t, ts, http.MethodPost, "/value/", strings.NewReader(jsonGetCounter))
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, jsonDoubleCounter, body)

	statusCode, body = testRequest(t, ts, http.MethodPost, "/value/", strings.NewReader(jsonGetGauge))
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, jsonGauge, body)
}

var (
	jsonCounter       = `{"id":"testCounter33","type":"counter","delta":5}`
	jsonDoubleCounter = `{"id":"testCounter33","type":"counter","delta":10}`
	jsonGauge         = `{"id":"testGauge33","type":"gauge","value":0.00001}`
	jsonGetCounter    = `{"id":"testCounter33","type":"counter"}`
	jsonGetGauge      = `{"id":"testGauge33","type":"gauge"}`
)

var htmlNoMetrics = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Метрики</title>
	</head>
	<body>
		<ul>
			<div><strong>no metrics</strong></div>
		</ul>
	</body>
</html>`

var htmlWithMetrics = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Метрики</title>
	</head>
	<body>
		<ul>
			<li><b>testCounter1:</b> 100</li><li><b>testGauge1:</b> 0.000001</li>
		</ul>
	</body>
</html>`
