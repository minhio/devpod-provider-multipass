package cmd

import (
	"github.com/minhio/devpod-provider-multipass/pkg/devpod"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop an instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return devpod.Stop()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
