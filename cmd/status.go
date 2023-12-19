package cmd

import (
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of an instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return multipass.Status()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
