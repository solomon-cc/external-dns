package dns

import "external-dns/pkg/sdk/tencent-cloud-sdk/sdk"

// Client is the sdk client struct, each func corresponds to an OpenAPI
type Client struct {
	sdk.Client
}

// NewClientWithAccessKey is a shortcut to create sdk client with accesskey
func NewClientWithSecretKey(accessKeyId, accessKeySecret string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithSecretKey(accessKeyId, accessKeySecret)
	return
}
