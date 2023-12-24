package cmd

import (
	"github.com/minhio/devpod-provider-multipass/pkg/devpod"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return devpod.Create()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
