package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rusalexch/metal/internal/services"
	"github.com/rusalexch/metal/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string) (int, string) {
	fmt.Println(ts.URL + path)
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	defer resp.Body.Close()

	return resp.StatusCode, string(respBody)
}

func TestNew(t *testing.T) {
	h := New(services.New(storage.New()))
	h.Init()
	ts := httptest.NewServer(h)
	defer ts.Close()

	statusCode, body := testRequest(t, ts, http.MethodGet, "/ping")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "pong", body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, htmlNoMetrics, body)

	statusCode, _ = testRequest(t, ts, http.MethodPost, "/update/gauge/testGauge1/1E-10")
	assert.Equal(t, http.StatusOK, statusCode)

	statusCode, body = testRequest(t, ts, http.MethodPost, "/update/unknown/testGauge1/1E-10")
	assert.Equal(t, http.StatusNotImplemented, statusCode)
	assert.Equal(t, "method not implemented", body)

	statusCode, _ = testRequest(t, ts, http.MethodPost, "/update/counter/testCounter1/100")
	assert.Equal(t, http.StatusOK, statusCode)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, htmlWithMetrics, body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/value/gauge/testGauge1")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "1E-10", body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/value/counter/testCounter1")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "100", body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/value/gauge/testGauge11")
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Equal(t, "metric not found", body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/value/counter/testCounter11")
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Equal(t, "metric not found", body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/value/noname/testCounter11")
	assert.Equal(t, http.StatusNotImplemented, statusCode)
	assert.Equal(t, "method not implemented", body)
}

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
			<li><b>testCounter1:</b> 100</li><li><b>testGauge1:</b> 1E-10</li>
		</ul>
	</body>
</html>`
