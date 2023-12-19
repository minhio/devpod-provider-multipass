/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start an instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return multipass.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
