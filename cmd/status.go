package cmd

import (
	"github.com/minhio/devpod-provider-multipass/pkg/devpod"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of an instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return devpod.Status()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
