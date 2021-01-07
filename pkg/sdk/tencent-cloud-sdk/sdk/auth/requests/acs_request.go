package requests

import (
	"io"
	"time"
)

// interface
type AcsRequest interface {
	GetScheme() string
	GetMethod() string
	GetDomain() string
	GetPort() string
	GetRegionId() string
	GetHeaders() map[string]string
	GetQueryParams() map[string]string
	GetFormParams() map[string]string
	GetContent() []byte
	GetBodyReader() io.Reader
	GetStyle() string
	GetProduct() string
	GetVersion() string
	GetActionName() string
	GetAcceptFormat() string
	GetLocationServiceCode() string
	GetLocationEndpointType() string
	GetReadTimeout() time.Duration
	GetConnectTimeout() time.Duration
	SetReadTimeout(readTimeout time.Duration)
	SetConnectTimeout(connectTimeout time.Duration)
	SetHTTPSInsecure(isInsecure bool)
	GetHTTPSInsecure() *bool

	GetUserAgent() map[string]string

	SetStringToSign(stringToSign string)
	GetStringToSign() string

	SetDomain(domain string)
	SetContent(content []byte)
	SetScheme(scheme string)
	BuildUrl() string
	BuildQueries() string

	addHeaderParam(key, value string)
	addQueryParam(key, value string)
	addFormParam(key, value string)
	addPathParam(key, value string)
}
