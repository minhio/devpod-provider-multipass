package cmd

import (
	"os"
	"os/exec"

	"github.com/loft-sh/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "devpod-provider-multipass",
	Short:         "multipass provider commands",
	SilenceErrors: true,
	SilenceUsage:  true,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		log.Default.MakeRaw()
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// execute command
	err := rootCmd.Execute()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if len(exitErr.Stderr) > 0 {
				log.Default.ErrorStreamOnly().Error(string(exitErr.Stderr))
			}
			os.Exit(exitErr.ExitCode())
		}

		log.Default.Fatal(err)
	}
}
