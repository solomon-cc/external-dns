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
	//Config   = &types.Domain{}
	cp providers.Provider
)

func init() {
	createCmd.Flags().StringVarP(&cProvider, "provider", "p", "", "Provider is a module which provides an interface for managing cloud resources")
	//createCmd.Flags().StringVarP(&Config.Domain, "domain", "d", "", "Domain is Host")
	//createCmd.Flags().StringVarP(&Config.SubDomain, "sub_domain", "s", "", "SubDomain is Host record")
	//createCmd.Flags().StringVarP(&Config.Value, "value", "a", "", "Value is record value of domain")
}

func CreateCommand() *cobra.Command {
	pStr := common.FlagHackLookup("--provider")
	if pStr != "" {
		if reg, err := providers.GetProvider(pStr); err != nil {
			logrus.Fatalln(err)
		} else {
			cp = reg
		}

		createCmd.Flags().AddFlagSet(cp.GetCredentialFlags(createCmd))
		createCmd.Flags().AddFlagSet(cp.GetCreateFlags(createCmd))

	}

	createCmd.Run = func(cmd *cobra.Command, args []string) {
		if cProvider == "" {
			logrus.Fatalln("required flags(s) \"[provider]\" not set")
		}

		// must bind after dynamic provider flags loaded. --TODO
		//common.BindPFlags(cmd, cp)

		if err := cp.CreateDnsRecord(); err != nil {
			logrus.Errorln(err)
		}

	}

	return createCmd
}
