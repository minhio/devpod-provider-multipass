package multipass

import (
	"os"
	"os/exec"

	"github.com/minhio/devpod-provider-multipass/pkg/options"
)

func Command() error {
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