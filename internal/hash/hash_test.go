package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"log"
	"testing"

	"github.com/rusalexch/metal/internal/app"
	"github.com/stretchr/testify/assert"
)

func wantHash(key, id string, delta *int64, value *float64) string {
	h := hmac.New(sha256.New, []byte(key))
	s := ""
	if delta != nil {
		s = fmt.Sprintf("%s:counter:%d", id, *delta)
	} else if value != nil {
		s = fmt.Sprintf("%s:gauge:%f", id, *value)
	}
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func TestNew(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want *Hash
	}{
		{
			name: "with empty key",
			args: args{key: ""},
			want: &Hash{hmac.New(sha256.New, []byte("")), true},
		},
		{
			name: "with key",
			args: args{key: "test"},
			want: &Hash{hmac.New(sha256.New, []byte("test")), false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.key)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHash_createHash(t *testing.T) {
	type args struct {
		m   app.Metrics
		key string
	}
	var counter int64 = 54
	gauge := 0.00001

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hash for counter",
			args: args{
				m: app.Metrics{
					ID:    "TestCounter",
					Type:  app.Counter,
					Delta: &counter,
					Value: nil,
				},
				key: "testKEY",
			},
			want: wantHash("testKEY", "TestCounter", &counter, nil),
		},
		{
			name: "hash for gauge",
			args: args{
				m: app.Metrics{
					ID:    "TestGuage",
					Type:  app.Gauge,
					Delta: nil,
					Value: &gauge,
				},
				key: "testKEY",
			},
			want: wantHash("testKEY", "TestGuage", nil, &gauge),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.args.key)
			got := h.createHash(tt.args.m)
			log.Println(got)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHash_AddHash(t *testing.T) {
	type args struct {
		m   app.Metrics
		key string
	}
	var counter int64 = 54

	tests := []struct {
		name   string
		args   args
		isHash bool
	}{
		{
			name: "with empty key",
			args: args{
				m: app.Metrics{
					ID:    "TestCounter",
					Type:  app.Counter,
					Delta: &counter,
					Value: nil,
				},
				key: "",
			},
			isHash: false,
		},
		{
			name: "with key",
			args: args{
				m: app.Metrics{
					ID:    "TestCounter",
					Type:  app.Counter,
					Delta: &counter,
					Value: nil,
				},
				key: "testKey",
			},
			isHash: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.args.key)
			h.AddHash(&tt.args.m)

			assert.Equal(t, tt.isHash, tt.args.m.Hash != "")
		})
	}
}

func TestHash_Check(t *testing.T) {
	type args struct {
		m   app.Metrics
		key string
	}
	gauge := 0.0000054

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "with empty key",
			args: args{
				m: app.Metrics{
					ID:    "TestCounter",
					Type:  app.Gauge,
					Delta: nil,
					Value: &gauge,
					Hash:  "",
				},
				key: "",
			},
			want: true,
		},
		{
			name: "with key",
			args: args{
				m: app.Metrics{
					ID:    "TestCounter",
					Type:  app.Gauge,
					Delta: nil,
					Value: &gauge,
					Hash:  wantHash("testKey", "TestCounter", nil, &gauge),
				},
				key: "testKey",
			},
			want: true,
		},
		{
			name: "with wrong key",
			args: args{
				m: app.Metrics{
					ID:    "TestCounter",
					Type:  app.Gauge,
					Delta: nil,
					Value: &gauge,
					Hash:  wantHash("testKey", "TestCounter", nil, &gauge),
				},
				key: "wrongTestKey",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.args.key)
			got := h.Check(tt.args.m)

			assert.Equal(t, tt.want, got)
		})
	}
}
