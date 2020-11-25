package tencent

import (
	"fmt"

	"external-dns/pkg/types"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)


func (p *Tencent) GetCreateFlags(cmd *cobra.Command) *pflag.FlagSet {
	fs := p.sharedFlags()

	for _, f := range fs {
		if f.ShortHand == "" {
			if cmd.Flags().Lookup(f.Name) == nil {
				switch t := f.V.(type) {
				case bool:
					cmd.Flags().BoolVar(f.P.(*bool), f.Name, t, f.Usage)
				case string:
					cmd.Flags().StringVar(f.P.(*string), f.Name, t, f.Usage)
				}
			}
		} else {
			if cmd.Flags().Lookup(f.Name) == nil {
				switch t := f.V.(type) {
				case bool:
					cmd.Flags().BoolVarP(f.P.(*bool), f.Name, f.ShortHand, t, f.Usage)
				case string:
					cmd.Flags().StringVarP(f.P.(*string), f.Name, f.ShortHand, t, f.Usage)
				}
			}
		}
	}

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		errFlags := make([]string, 0)
		for _, f := range fs {
			if f.Required {
				p, ok := f.P.(*string)
				if ok {
					if *p == "" && f.V.(string) == "" {
						errFlags = append(errFlags, f.Name)
					}
				}
			}
		}

		if len(errFlags) == 0 {
			return nil
		}

		return fmt.Errorf("required flags(s) \"%s\" not set", errFlags)
	}

	return cmd.Flags()
}

func (p *Tencent) GetCredentialFlags(cmd *cobra.Command) *pflag.FlagSet {
	fs := []types.Flag{
		{
			Name:     secretID,
			P:        &p.SecretID,
			V:        p.SecretID,
			Usage:    "User access key ID",
			Required: true,
		},
		{
			Name:     secretKey,
			P:        &p.SecretKey,
			V:        p.SecretKey,
			Usage:    "User access key secret",
			Required: true,
		},
	}

	for _, f := range fs {
		if f.ShortHand == "" {
			if cmd.Flags().Lookup(f.Name) == nil {
				switch t := f.V.(type) {
				case bool:
					cmd.Flags().BoolVar(f.P.(*bool), f.Name, t, f.Usage)
				case string:
					cmd.Flags().StringVar(f.P.(*string), f.Name, t, f.Usage)
				}
			}
		} else {
			if cmd.Flags().Lookup(f.Name) == nil {
				switch t := f.V.(type) {
				case bool:
					cmd.Flags().BoolVarP(f.P.(*bool), f.Name, f.ShortHand, t, f.Usage)
				case string:
					cmd.Flags().StringVarP(f.P.(*string), f.Name, f.ShortHand, t, f.Usage)
				}
			}
		}
	}

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		errFlags := make([]string, 0)
		for _, f := range fs {
			if f.Required {
				p, ok := f.P.(*string)
				if ok {
					if *p == "" && f.V.(string) == "" {
						errFlags = append(errFlags, f.Name)
					}
				}
			}
		}
		if len(errFlags) == 0 {
			return nil
		}

		return fmt.Errorf("required flags(s) \"%s\" not set", errFlags)
	}

	return cmd.Flags()
}

func (p *Tencent) BindCredentialFlags() *pflag.FlagSet {
	nfs := pflag.NewFlagSet("", pflag.ContinueOnError)
	nfs.StringVar(&p.SecretID, secretID, p.SecretID, "User access key ID")
	nfs.StringVar(&p.SecretKey, secretKey, p.SecretKey, "User access key secret")
	return nfs
}

func (p *Tencent) sharedFlags() []types.Flag {
	fs := []types.Flag{
		{
			Name:     "domain",
			P:        &p.Domain,
			V:        p.Domain,
			Usage:    "Domain is Host",
			Required: true,
		},
		{
			Name:     "sub-domain",
			P:        &p.SubDomain,
			V:        p.SubDomain,
			Usage:    "SubDomain is Host record",
			Required: true,
		},
		{
			Name:     "value",
			P:        &p.Value,
			V:        p.Value,
			Usage:    "Value is record of domain",
			Required: true,
		},
	}

	return fs
}