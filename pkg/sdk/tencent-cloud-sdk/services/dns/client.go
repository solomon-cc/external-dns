package dns

import "external-dns/pkg/sdk/tencent-cloud-sdk/sdk"

// Client is the sdk client struct, each func corresponds to an OpenAPI
type Client struct {
	sdk.Client
}

// NewClientWithAccessKey is a shortcut to create sdk client with accesskey
// usage: https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md
func NewClientWithAccessKey(accessKeyId, accessKeySecret string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithAccessKey(accessKeyId, accessKeySecret)
	return
}
