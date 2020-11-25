package types

type Flag struct {
	Name      string
	P         interface{}
	V         interface{}
	ShortHand string
	Usage     string
	Required  bool
}

type Domain struct {
	Domain    string `json:"domain,omitempty" yaml:"domain,omitempty"`
	SubDomain string `json:"sub-domain,omitempty" yaml:"sub-domain,omitempty"`
	Value     string `json:"value,omitempty" yaml:"value,omitempty"`
}
