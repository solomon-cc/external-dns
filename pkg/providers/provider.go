package providers

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Factory is a function that returns a Provider.Interface.
type Factory func() (Provider, error)

var (
	providersMutex sync.Mutex
	providers      = make(map[string]Factory)
)

// Provider is an abstract, pluggable interface for dns provider
type Provider interface {
	GetProviderName() string
	// Create command flags.
	GetCreateFlags(cmd *cobra.Command) *pflag.FlagSet
	// Credential flags.
	GetCredentialFlags(cmd *cobra.Command) *pflag.FlagSet
	// Use this method to bind Viper, although it is somewhat repetitive.
	BindCredentialFlags() *pflag.FlagSet
	// Create DNS record
	CreateDnsRecord() error
	// Update DNS record
	//UpdateDnsRecord() error
	// Delete DNS record
	//DeleteDnsRecord() error
}

// RegisterProvider registers a provider.Factory by name.
func RegisterProvider(name string, p Factory) {
	providersMutex.Lock()
	defer providersMutex.Unlock()
	if _, found := providers[name]; !found {
		logrus.Debugf("registered provider %s", name)
		providers[name] = p
	}
}

// GetProvider creates an instance of the named provider, or nil if
// the name is unknown.  The error return is only used if the named provider
// was known but failed to initialize.
func GetProvider(name string) (Provider, error) {
	providersMutex.Lock()
	defer providersMutex.Unlock()
	f, found := providers[name]
	if !found {
		return nil, fmt.Errorf("provider %s is not registered", name)
	}
	return f()
}
