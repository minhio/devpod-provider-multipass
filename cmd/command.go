package cmd

import (
	"github.com/minhio/devpod-provider-multipass/pkg/devpod"
	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Command an instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return devpod.Command()
	},
}

func init() {
	rootCmd.AddCommand(commandCmd)
}
