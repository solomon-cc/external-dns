package sdk

import (
	"net/http"
	"sync"
	"time"
)

// Client the type Client
type Client struct {
	isInsecure     bool
	regionId       string
	httpProxy      string
	httpsProxy     string
	noProxy        string
	userAgent      map[string]string
	httpClient     *http.Client
	asyncTaskQueue chan func()
	readTimeout    time.Duration
	connectTimeout time.Duration
	EndpointMap    map[string]string
	EndpointType   string
	Network        string
	Domain         string

	debug     bool
	isRunning bool
	// void "panic(write to close channel)" cause of addAsync() after Shutdown()
	asyncChanLock *sync.RWMutex
}
