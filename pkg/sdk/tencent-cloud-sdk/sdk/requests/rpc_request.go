package requests

import (
	"fmt"
	"io"
	"strings"

	"external-dns/pkg/sdk/tencent-cloud-sdk/sdk/utils"
)

type RpcRequest struct {
	*BaseRequest
}

func (request *RpcRequest) init() {
	request.BaseRequest = defaultBaseRequest()
	request.Method = POST
}

func (*RpcRequest) GetStyle() string {
	return RPC
}

func (request *RpcRequest) GetBodyReader() io.Reader {
	if request.FormParams != nil && len(request.FormParams) > 0 {
		formString := utils.GetUrlFormedMap(request.FormParams)
		return strings.NewReader(formString)
	} else {
		return strings.NewReader("")
	}
}

func (request *RpcRequest) BuildQueries() string {
	request.queries = "/?" + utils.GetUrlFormedMap(request.QueryParams)
	return request.queries
}

func (request *RpcRequest) BuildUrl() string {
	url := fmt.Sprintf("%s://%s", strings.ToLower(request.Scheme), request.Domain)
	if len(request.Port) > 0 {
		url = fmt.Sprintf("%s:%s", url, request.Port)
	}
	return url + request.BuildQueries()
}

func (request *RpcRequest) GetVersion() string {
	return request.version
}

func (request *RpcRequest) GetActionName() string {
	return request.actionName
}

func (request *RpcRequest) addPathParam(key, value string) {
	panic("not support")
}

func (request *RpcRequest) InitWithApiInfo(product, version, action, serviceCode, endpointType string) {
	request.init()
	request.product = product
	request.version = version
	request.actionName = action
	request.locationServiceCode = serviceCode
	request.locationEndpointType = endpointType
}
