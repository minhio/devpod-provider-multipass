package multipass

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/minhio/devpod-provider-multipass/pkg/options"
)

func Command() error {
	devPodCommand := os.Getenv("COMMAND")
	if devPodCommand == "" {
		return fmt.Errorf("command environment variable is missing")
	}

	multipassOptions, err := options.FromEnv()
	if err != nil {
		return err
	}

	machineId := multipassOptions.GetMachineId()

	cmd := exec.Command(multipassOptions.Path, "exec", machineId, "--", devPodCommand)
	cmd.Env = os.Environ()
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func Create() error {
	return nil
}
func Delete() error {
	return nil
}

func Init() error {
	multipassOptions, err := options.FromEnv()
	if err != nil {
		return err
	}

	// execute 'multipass version' command
	// as a way to check if multipass is available
	cmd := exec.Command(multipassOptions.Path, "version")
	cmd.Env = os.Environ()
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func Start() error {
	return nil
}

func Status() error {
	return nil
}

func Stop() error {
	return nil
}
