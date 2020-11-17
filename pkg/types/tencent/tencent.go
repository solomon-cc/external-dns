package tencent

type Options struct {
	AccessKey    string `json:"access-key,omitempty" yaml:"access-key,omitempty"`
	AccessSecret string `json:"access-secret,omitempty" yaml:"access-secret,omitempty"`
	Action       string `json:"action,omitempty" yaml:"action,omitempty"`
	Domain       string `json:"domain,omitempty" yaml:"domain,omitempty"`
	Nonce        string `json:"nonce,omitempty" yaml:"nonce,omitempty"`
	Offset       string `json:"offset,omitempty" yaml:"offset,omitempty"`
	SubDomain    string `json:"sub-domain,omitempty" yaml:"sub-domain,omitempty"`
	Timestamp    string `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
	Value        string `json:"value,omitempty" yaml:"value,omitempty"`
}
