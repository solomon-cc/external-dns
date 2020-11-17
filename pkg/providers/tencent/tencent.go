package tencent

import (
	"sync"

	"external-dns/pkg/types/tencent"

	"github.com/sirupsen/logrus"
)

// ProviderName is the name of this provider.
const ProviderName = "tencent"

type checkFun func() error

type Tencent struct {
	tencent.Options `json:",inline"`

	m      *sync.Map
	logger *logrus.Logger

}


func (p *Tencent) GetProviderName() string {
	return ProviderName
}

func (p *Tencent) generateClientSDK() error {

	return nil
}