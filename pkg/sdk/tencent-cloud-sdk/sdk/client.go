package sdk

import (
	"net/http"
	"sync"
	"time"

	"external-dns/pkg/sdk/tencent-cloud-sdk/sdk/auth/credentials"
)

// Client the type Client
type Client struct {
	httpClient     *http.Client
	config         *Config
	asyncTaskQueue chan func()
	readTimeout    time.Duration
	connectTimeout time.Duration
	Domain         string

	debug     bool
	isRunning bool
	// void "panic(write to close channel)" cause of addAsync() after Shutdown()
	asyncChanLock *sync.RWMutex
}

func (client *Client) Init() (err error) {
	panic("not support yet")
}

func NewClient() (client *Client, err error) {
	client = &Client{}
	err = client.Init()
	return
}

func (client *Client) InitWithSecretKey(secretID, secretKey string) (err error) {
	config := client.InitClientConfig()
	_ = &credentials.AccessKeyCredential{
		SecretID:  secretID,
		SecretKey: secretKey,
	}
	return client.InitWithOptions(config)
}

// EnableAsync enable the async task queue
func (client *Client) EnableAsync(routinePoolSize, maxTaskQueueSize int) {
	client.asyncTaskQueue = make(chan func(), maxTaskQueueSize)
	for i := 0; i < routinePoolSize; i++ {
		go func() {
			for client.isRunning {
				select {
				case task, notClosed := <-client.asyncTaskQueue:
					if notClosed {
						task()
					}
				}
			}
		}()
	}
}

func (client *Client) InitClientConfig() (config *Config) {
	if client.config != nil {
		return client.config
	} else {
		return NewConfig()
	}
}

func (client *Client) InitWithOptions(config *Config) (err error) {
	client.isRunning = true
	client.asyncChanLock = new(sync.RWMutex)
	client.config = config
	client.httpClient = &http.Client{}

	if config.Transport != nil {
		client.httpClient.Transport = config.Transport
	} else if config.HttpTransport != nil {
		client.httpClient.Transport = config.HttpTransport
	}

	if config.Timeout > 0 {
		client.httpClient.Timeout = config.Timeout
	}

	if config.EnableAsync {
		client.EnableAsync(config.GoRoutinePoolSize, config.MaxTaskQueueSize)
	}

	return
}

// TODO
func (client *Client) DoAction() (err error) {
	return
}

// TODO
func (client *Client) DoActionWithSigner() (err error) {
	return
}
