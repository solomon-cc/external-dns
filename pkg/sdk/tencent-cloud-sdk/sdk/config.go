package sdk

import (
	"net/http"
	"time"
)

type Config struct {
	AutoRetry         bool              `default:"true"`
	MaxRetryTime      int               `default:"3"`
	UserAgent         string            `default:""`
	Debug             bool              `default:"false"`
	HttpTransport     *http.Transport   `default:""`
	Transport         http.RoundTripper `default:""`
	EnableAsync       bool              `default:"false"`
	MaxTaskQueueSize  int               `default:"1000"`
	GoRoutinePoolSize int               `default:"5"`
	Scheme            string            `default:"HTTP"`
	Timeout           time.Duration
}

func NewConfig() (config *Config) {
	config = &Config{}
	return
}