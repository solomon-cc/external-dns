package cmd

import (
	"external-dns/cmd/common"
	"external-dns/pkg/providers"
	_ "external-dns/pkg/providers/tencent"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create dns record",
		Example: `  external-dns create \
    --provider tencent \
    --domain <domain name> \
    --sub_domain <sub_domain name> \
    --value <value>`,
	}

	cProvider = ""
	cDomain = ""
	cSubDomain = ""
	cValue = ""
	cp        providers.Provider
)

func init() {
	createCmd.Flags().StringVarP(&cProvider, "provider", "p", cProvider, "Provider is a module which provides an interface for managing cloud resources")
	createCmd.Flags().StringVarP(&cDomain, "domain", "d", cDomain, "Domain is Host")
	createCmd.Flags().StringVarP(&cSubDomain, "sub_domain", "s", cSubDomain, "SubDomain is Host record")
	createCmd.Flags().StringVarP(&cValue, "value", "a", cValue, "Value is record value of domain")
}

func CreateCommand() *cobra.Command {
	pStr := common.FlagHackLookup("--provider")
	if pStr != "" {
		if reg, err := providers.GetProvider(pStr); err != nil {
			logrus.Fatalln(err)
		} else {
			cp = reg
		}
	}

	createCmd.Run = func(cmd *cobra.Command, args []string) {
		if cProvider == "" {
			logrus.Fatalln("required flags(s) \"[provider]\" not set")
		}

	}

	return createCmd
}
