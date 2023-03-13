package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rusalexch/metal/internal/services"
	"github.com/rusalexch/metal/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlers_update(t *testing.T) {
	type want struct {
		code int
		res  string
	}
	type args struct {
		url         string
		method      string
		contentType string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "should be ok",
			args: args{
				url:         "/update/gauge/test/123.32",
				method:      http.MethodPost,
				contentType: "text/plain",
			},
			want: want{
				code: http.StatusOK,
				res:  "",
			},
		},
		{
			name: "should be method error",
			args: args{
				url:         "/update/gauge/test/123.32",
				method:      http.MethodGet,
				contentType: "text/plain",
			},
			want: want{
				code: http.StatusBadRequest,
				res:  "method not available",
			},
		},
		{
			name: "should be error count params",
			args: args{
				url:         "/update/gauge/test",
				method:      http.MethodPost,
				contentType: "text/plain",
			},
			want: want{
				code: http.StatusNotFound,
				res:  "required three params",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := New(services.New(storage.New()))
			req := httptest.NewRequest(tt.args.method, tt.args.url, nil)
			req.Header.Add("Content-Type", tt.args.contentType)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(handler.update)
			h.ServeHTTP(w, req)
			res := w.Result()

			assert.Equal(t, tt.want.code, res.StatusCode)
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.Equal(t, tt.want.res, string(body))
		})
	}
}
