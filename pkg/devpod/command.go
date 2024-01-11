package devpod

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/loft-sh/devpod/pkg/ssh"
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
	"github.com/pkg/errors"
)

// devpod uses this to inject itself into the environment and
// route all communication through the commands standard output and input.
func Command() error {
	// devpod sets the command to be executed as an env var
	// here we are retrieving the command to be executed
	devPodCommand := os.Getenv("COMMAND")
	if devPodCommand == "" {
		return errors.New("command environment variable is missing")
	}

	log.Default().Printf("[devpod] command: %s", devPodCommand)

	// get multipass options from env vars
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	// build the machine context from env vars
	machine := provider.FromEnvironment()

	// init multipass client
	client, err := multipass.NewClient(opts.Path)
	if err != nil {
		return err
	}

	// get multipass instance info by machine ID (the instance name)
	instance, err := client.GetInstance(machine.ID)
	if err != nil {
		return err
	}

	// get the instance's ssh private key
	privateKey, err := ssh.GetPrivateKeyRawBase(machine.Folder)
	if err != nil {
		return fmt.Errorf("load private key: %w", err)
	}

	// build ssh client
	sshClient, err := ssh.NewSSHClient("devpod", instance.Ipv4[0]+":22", privateKey)
	if err != nil {
		return errors.Wrap(err, "create ssh client")
	}
	defer sshClient.Close()

	// run command
	return ssh.Run(context.Background(), sshClient, devPodCommand, os.Stdin, os.Stdout, os.Stderr)
}
