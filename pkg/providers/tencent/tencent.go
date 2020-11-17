package tencent

import "github.com/sirupsen/logrus"

// ProviderName is the name of this provider.
const ProviderName = "tencent"

type checkFun func() error

type Tencent struct {


	logger *logrus.Logger
}


func (p *Tencent) GetProviderName() string {
	return ProviderName
}