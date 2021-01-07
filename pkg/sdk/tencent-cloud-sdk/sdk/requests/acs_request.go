package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"

	"external-dns/pkg/sdk/tencent-cloud-sdk/sdk/errors"
)

const (
	RPC = "RPC"
	ROA = "ROA"

	HTTP  = "HTTP"
	HTTPS = "HTTPS"

	DefaultHttpPort = "80"

	GET     = "GET"
	PUT     = "PUT"
	POST    = "POST"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"

	Json = "application/json"
	Xml  = "application/xml"
	Raw  = "application/octet-stream"
	Form = "application/x-www-form-urlencoded"

	Header = "Header"
	Query  = "Query"
	Body   = "Body"
	Path   = "Path"

	HeaderSeparator = "\n"
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

// base class
type BaseRequest struct {
	Scheme         string
	Method         string
	Domain         string
	Port           string
	RegionId       string
	ReadTimeout    time.Duration
	ConnectTimeout time.Duration
	isInsecure     *bool

	userAgent map[string]string
	product   string
	version   string

	actionName string

	AcceptFormat string

	QueryParams map[string]string
	Headers     map[string]string
	FormParams  map[string]string
	Content     []byte

	locationServiceCode  string
	locationEndpointType string

	queries string

	stringToSign string
}

func (request *BaseRequest) GetQueryParams() map[string]string {
	return request.QueryParams
}

func (request *BaseRequest) GetFormParams() map[string]string {
	return request.FormParams
}

func (request *BaseRequest) GetReadTimeout() time.Duration {
	return request.ReadTimeout
}

func (request *BaseRequest) GetConnectTimeout() time.Duration {
	return request.ConnectTimeout
}

func (request *BaseRequest) SetReadTimeout(readTimeout time.Duration) {
	request.ReadTimeout = readTimeout
}

func (request *BaseRequest) SetConnectTimeout(connectTimeout time.Duration) {
	request.ConnectTimeout = connectTimeout
}

func (request *BaseRequest) GetHTTPSInsecure() *bool {
	return request.isInsecure
}

func (request *BaseRequest) SetHTTPSInsecure(isInsecure bool) {
	request.isInsecure = &isInsecure
}

func (request *BaseRequest) GetContent() []byte {
	return request.Content
}

func (request *BaseRequest) GetVersion() string {
	return request.version
}

func (request *BaseRequest) GetActionName() string {
	return request.actionName
}

func (request *BaseRequest) SetContent(content []byte) {
	request.Content = content
}

func (request *BaseRequest) GetUserAgent() map[string]string {
	return request.userAgent
}

func (request *BaseRequest) AppendUserAgent(key, value string) {
	newkey := true
	if request.userAgent == nil {
		request.userAgent = make(map[string]string)
	}
	if strings.ToLower(key) != "core" && strings.ToLower(key) != "go" {
		for tag, _ := range request.userAgent {
			if tag == key {
				request.userAgent[tag] = value
				newkey = false
			}
		}
		if newkey {
			request.userAgent[key] = value
		}
	}
}

func (request *BaseRequest) addHeaderParam(key, value string) {
	request.Headers[key] = value
}

func (request *BaseRequest) addQueryParam(key, value string) {
	request.QueryParams[key] = value
}

func (request *BaseRequest) addFormParam(key, value string) {
	request.FormParams[key] = value
}

func (request *BaseRequest) GetAcceptFormat() string {
	return request.AcceptFormat
}

func (request *BaseRequest) GetLocationServiceCode() string {
	return request.locationServiceCode
}

func (request *BaseRequest) GetLocationEndpointType() string {
	return request.locationEndpointType
}

func (request *BaseRequest) GetProduct() string {
	return request.product
}

func (request *BaseRequest) GetScheme() string {
	return request.Scheme
}

func (request *BaseRequest) SetScheme(scheme string) {
	request.Scheme = scheme
}

func (request *BaseRequest) GetMethod() string {
	return request.Method
}

func (request *BaseRequest) GetDomain() string {
	return request.Domain
}

func (request *BaseRequest) SetDomain(host string) {
	request.Domain = host
}

func (request *BaseRequest) GetPort() string {
	return request.Port
}

func (request *BaseRequest) GetRegionId() string {
	return request.RegionId
}

func (request *BaseRequest) GetHeaders() map[string]string {
	return request.Headers
}

func (request *BaseRequest) SetContentType(contentType string) {
	request.addHeaderParam("Content-Type", contentType)
}

func (request *BaseRequest) GetContentType() (contentType string, contains bool) {
	contentType, contains = request.Headers["Content-Type"]
	return
}

func (request *BaseRequest) SetStringToSign(stringToSign string) {
	request.stringToSign = stringToSign
}

func (request *BaseRequest) GetStringToSign() string {
	return request.stringToSign
}

func defaultBaseRequest() (request *BaseRequest) {
	request = &BaseRequest{
		Scheme:       "",
		AcceptFormat: "JSON",
		Method:       GET,
		QueryParams:  make(map[string]string),
		Headers: map[string]string{
			"x-sdk-client":      "golang/1.0.0",
			"x-sdk-invoke-type": "normal",
			"Accept-Encoding":   "identity",
		},
		FormParams: make(map[string]string),
	}
	return
}

func InitParams(request AcsRequest) (err error) {
	requestValue := reflect.ValueOf(request).Elem()
	err = flatRepeatedList(requestValue, request, "", "")
	return
}

func flatRepeatedList(dataValue reflect.Value, request AcsRequest, position, prefix string) (err error) {
	dataType := dataValue.Type()
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		name, containsNameTag := field.Tag.Lookup("name")
		fieldPosition := position
		if fieldPosition == "" {
			fieldPosition, _ = field.Tag.Lookup("position")
		}
		typeTag, containsTypeTag := field.Tag.Lookup("type")
		if containsNameTag {
			if !containsTypeTag {
				// simple param
				key := prefix + name
				value := dataValue.Field(i).String()
				if dataValue.Field(i).Kind().String() == "map" {
					byt, _ := json.Marshal(dataValue.Field(i).Interface())
					value = string(byt)
				}
				err = addParam(request, fieldPosition, key, value)
				if err != nil {
					return
				}
			} else if typeTag == "Repeated" {
				// repeated param
				repeatedFieldValue := dataValue.Field(i)
				if repeatedFieldValue.Kind() != reflect.Slice {
					// possible value: {"[]string", "*[]struct"}, we must call Elem() in the last condition
					repeatedFieldValue = repeatedFieldValue.Elem()
				}
				if repeatedFieldValue.IsValid() && !repeatedFieldValue.IsNil() {
					for m := 0; m < repeatedFieldValue.Len(); m++ {
						elementValue := repeatedFieldValue.Index(m)
						key := prefix + name + "." + strconv.Itoa(m+1)
						if elementValue.Type().Kind().String() == "string" {
							value := elementValue.String()
							err = addParam(request, fieldPosition, key, value)
							if err != nil {
								return
							}
						} else {
							err = flatRepeatedList(elementValue, request, fieldPosition, key+".")
							if err != nil {
								return
							}
						}
					}
				}
			}
		}
	}
	return
}

func addParam(request AcsRequest, position, name, value string) (err error) {
	if len(value) > 0 {
		switch position {
		case Header:
			request.addHeaderParam(name, value)
		case Query:
			request.addQueryParam(name, value)
		case Path:
			request.addPathParam(name, value)
		case Body:
			request.addFormParam(name, value)
		default:
			errMsg := fmt.Sprintf(errors.UnsupportedParamPositionErrorMessage, position)
			err = errors.NewClientError(errors.UnsupportedParamPositionErrorCode, errMsg, nil)
		}
	}
	return
}