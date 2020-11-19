package main

import (
	"os"

	"external-dns/cmd"
)

var (
	gitVersion   string
	gitCommit    string
	gitTreeState string
	buildDate    string
)


func main() {
	rootCmd := cmd.Command()

	rootCmd.AddCommand(cmd.VersionCommand(gitVersion, gitCommit, gitTreeState, buildDate),cmd.CreateCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
