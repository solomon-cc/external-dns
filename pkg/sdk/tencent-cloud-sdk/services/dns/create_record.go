package dns

import (
	"external-dns/pkg/sdk/tencent-cloud-sdk/sdk/responses"
	"external-dns/pkg/types/tencent"
)

type AddDomainRecordRequest struct {
	*tencent.Options
}

func CreateAddDomainRecordRequest() (request *AddDomainRecordRequest) {
	request = &AddDomainRecordRequest{
		Options: &tencent.Options{
			TTL: 600,
		},
	}

	return
}

type AddDomainRecordResponse struct {
	*responses.BaseResponse
}
