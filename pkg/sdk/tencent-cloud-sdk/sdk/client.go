package sdk

import (
	"external-dns/pkg/sdk/tencent-cloud-sdk/sdk/requests"
	"external-dns/pkg/sdk/tencent-cloud-sdk/sdk/responses"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"external-dns/pkg/sdk/tencent-cloud-sdk/sdk/auth/credentials"
	"external-dns/pkg/sdk/tencent-cloud-sdk/sdk/auth/signers"
	"external-dns/pkg/sdk/tencent-cloud-sdk/services/dns"
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
func (client *Client) DoAction(request requests.AcsRequest, response responses.AcsResponse) (err error) {
	response,_ = client.httpClient.Get()
	return
}


func (client *Client) DoActionWithSigner(request *dns.AddDnsRecordRequest) (signStr string) {
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	rand.Seed(time.Now().UnixNano())
	x := strconv.Itoa(rand.Intn(65535310))

	uri := fmt.Sprintf("GETcns.api.qcloud.com/v2/index.php?"+
		"Action=%s"+
		"&Nonce=%s"+
		"&SecretId=%s"+
		"&SignatureMethod=HmacSHA256"+
		"&Timestamp=%s", request.Action, x, request.SecretID, timeStamp)

	params := fmt.Sprintf(
		"&Nonce=%s"+
			"&SecretId=%s"+
			"&Signature=%s"+
			"&SignatureMethod=HmacSHA256"+
			"&Timestamp=%s"+
			"&Action=%s"+
			"&domain=%s"+
			"&subDomain=%s"+
			"&recordType=A"+
			"&recodline=%s"+
			"&value=%s", x, request.SecretID, url.QueryEscape(signers.ComputeHmacSha256(uri, request.SecretKey)), timeStamp, request.Action, request.Domain, request.SubDomain, "默认", request.Value)

	return params
}
