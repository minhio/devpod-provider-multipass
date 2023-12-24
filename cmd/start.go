package cmd

import (
	"github.com/minhio/devpod-provider-multipass/pkg/devpod"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start an instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return devpod.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
