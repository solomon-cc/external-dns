package dns

import (
	"external-dns/pkg/sdk/tencent-cloud-sdk/sdk/requests"
	"external-dns/pkg/sdk/tencent-cloud-sdk/sdk/responses"
	"external-dns/pkg/types/tencent"
)

func (client *Client) AddDnsRecord(request *AddDnsRecordRequest) (response *AddDnsRecordResponse, err error) {
	response = CreateAddDnsRecordResponse()
	err = client.DoAction(request,response)
	return
}

type AddDnsRecordRequest struct {
	*requests.RpcRequest
	*tencent.Options
}

func CreateAddDomainRecordRequest() (request *AddDnsRecordRequest) {
	request = &AddDnsRecordRequest{
		Options: &tencent.Options{
			TTL: 600,
			Action: "RecordCreate",
		},
	}

	return
}

//TODO
func CreateAddDnsRecordResponse() (response *AddDnsRecordResponse) {
	response = &AddDnsRecordResponse{

	}

	return
}

type AddDnsRecordResponse struct {
	*responses.BaseResponse
}
