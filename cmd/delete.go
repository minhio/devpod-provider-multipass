package cmd

import (
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return multipass.Delete()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
