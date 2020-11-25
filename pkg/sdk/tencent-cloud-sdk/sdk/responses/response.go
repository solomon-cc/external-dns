package responses

import (
	"io/ioutil"
	"net/http"
)

type BaseResponse struct {
	httpStatus         int
	httpHeaders        map[string][]string
	httpContentString  string
	httpContentBytes   []byte
	originHttpResponse *http.Response
}


func (baseResponse *BaseResponse) GetHttpStatus() int {
	return baseResponse.httpStatus
}

func (baseResponse *BaseResponse) GetHttpHeaders() map[string][]string {
	return baseResponse.httpHeaders
}

func (baseResponse *BaseResponse) GetOriginHttpResponse() *http.Response {
	return baseResponse.originHttpResponse
}

func (baseResponse *BaseResponse) IsSuccess() bool {
	if baseResponse.GetHttpStatus() >= 200 && baseResponse.GetHttpStatus() < 300 {
		return true
	}

	return false
}

func (baseResponse *BaseResponse) parseFromHttpResponse(httpResponse *http.Response) (err error) {
	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}
	baseResponse.httpStatus = httpResponse.StatusCode
	baseResponse.httpHeaders = httpResponse.Header
	baseResponse.httpContentBytes = body
	baseResponse.httpContentString = string(body)
	baseResponse.originHttpResponse = httpResponse
	return
}
