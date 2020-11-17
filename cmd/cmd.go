package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:              "external-dns",
		Short:            "external-dns is used to manage the dns record on multiple cloud providers",
		Long:             `external-dns is used to manage the dns record on multiple cloud providers`,
		TraverseChildren: true,
	}
)

func Command() *cobra.Command {
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			logrus.Errorln(err)
			os.Exit(1)
		}
	}
	return cmd
}