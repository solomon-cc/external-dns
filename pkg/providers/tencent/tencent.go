package tencent

import (
	"sync"

	"external-dns/pkg/common"
	"external-dns/pkg/providers"
	"external-dns/pkg/sdk/tencent-cloud-sdk/services/dns"
	"external-dns/pkg/types/tencent"

	"github.com/sirupsen/logrus"
)

// ProviderName is the name of this provider.
const (
	ProviderName = "tencent"
	secretID     = "secret-id"
	secretKey    = "secret-key"
)

type checkFun func() error

type Tencent struct {
	tencent.Options `json:",inline"`

	s      *dns.Client
	m      *sync.Map
	logger *logrus.Logger
}

func init() {
	providers.RegisterProvider(ProviderName, func() (providers.Provider, error) {
		return NewProvider(), nil
	})
}

func NewProvider() *Tencent {
	return &Tencent{
		Options: tencent.Options{
			TTL: 600,
		},
		m: new(sync.Map),
	}

}

func (p *Tencent) GetProviderName() string {
	return ProviderName
}

func (p *Tencent) CreateDnsRecord() error {
	p.logger = common.NewLogger(common.Debug)
	p.logger.Infof("[%s] executing create dns record...\n", p.GetProviderName())

	if err := p.generateClientSDK(); err != nil {
		return err
	}

	request := dns.CreateAddDomainRecordRequest()
	logrus.Info(p)
	request.Value = p.Value
	request.Domain = p.Domain
	request.SubDomain = p.SubDomain

	return nil
}

func (p *Tencent) generateClientSDK() error {
	client, err := dns.NewClientWithSecretKey(p.SecretID, p.SecretKey)
	if err != nil {
		return err
	}
	client.EnableAsync(5, 1000)
	p.s = client

	return nil
}
