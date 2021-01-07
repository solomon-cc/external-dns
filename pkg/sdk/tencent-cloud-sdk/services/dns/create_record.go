package dns

import (
	"external-dns/pkg/sdk/tencent-cloud-sdk/sdk/responses"
	"external-dns/pkg/types/tencent"
)

func (client *Client) AddDnsRecord(request *AddDnsRecordRequest) (response *AddDnsRecordResponse, err error) {
	response = CreateAddDnsRecordResponse()
	err = client.DoAction(request)
	return
}

type AddDnsRecordRequest struct {
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
