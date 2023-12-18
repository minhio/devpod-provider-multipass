package cmd

import (
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init multipass",
	RunE: func(cmd *cobra.Command, args []string) error {
		return multipass.Init()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
