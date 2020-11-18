package tencent

type Options struct {
	Domain       string `json:"domain,omitempty" yaml:"domain,omitempty"`
	SubDomain    string `json:"sub-domain,omitempty" yaml:"sub-domain,omitempty"`
	RecordLine   string `json:"record_line,omitempty" yaml:"record_line,omitempty"`
	RecordLineId string `json:"record_line_id,omitempty" yaml:"record_line_id,omitempty"`
	LoginToken   string `json:"login_token,omitempty" yaml:"login_token,omitempty"`
	MX           string `json:"mx,omitempty" yaml:"mx,omitempty"`
	TTL          string `json:"ttl,omitempty" yaml:"ttl,omitempty"`
	Value        string `json:"value,omitempty" yaml:"value,omitempty"`
}
