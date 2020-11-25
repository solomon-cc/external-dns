package tencent

type Options struct {
	Action       string `json:"action,omitempty" yaml:"action,omitempty"`
	SecretID    string `json:"secret-id,omitempty" yaml:"secret-id,omitempty"`
	SecretKey string `json:"secret-key,omitempty" yaml:"secret-key,omitempty"`
	Domain       string `json:"domain,omitempty" yaml:"domain,omitempty"`
	Nonce        string `json:"nonce,omitempty" yaml:"nonce,omitempty"`
	SubDomain    string `json:"sub-domain,omitempty" yaml:"sub-domain,omitempty"`
	Timestamp    string `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
	TTL          int    `json:"ttl,omitempty" yaml:"ttl,omitempty"`
	Value        string `json:"value,omitempty" yaml:"value,omitempty"`
}

type CloudControllerManager struct {
	Region       string `json:"region,omitempty" yaml:"region,omitempty"`
	SecretID    string `json:"secret-id,omitempty" yaml:"secret-id,omitempty"`
	SecretKey string `json:"secret-key,omitempty" yaml:"secret-key,omitempty"`
}
