package cmd

import (
	"github.com/minhio/devpod-provider-multipass/pkg/devpod"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return devpod.Delete()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
