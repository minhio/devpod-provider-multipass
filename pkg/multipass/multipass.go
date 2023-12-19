package multipass

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/minhio/devpod-provider-multipass/pkg/options"
)

func Command() error {
	devPodCommand := os.Getenv("COMMAND")
	if devPodCommand == "" {
		return fmt.Errorf("command environment variable is missing")
	}

	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machineId := opts.GetMachineId()

	cmd := exec.Command(opts.Path, "exec", machineId, "--", devPodCommand)
	cmd.Env = os.Environ()
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func Create() error {
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machineId := opts.GetMachineId()

	cmd := exec.Command(opts.Path, "launch",
		"--image", opts.Image,
		"--cpus", strconv.Itoa(opts.Cpus),
		"--disk", opts.DiskSize,
		"--memory", opts.Memory,
		"--name", machineId,
		opts.Image,
	)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func Delete() error {
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machineId := opts.GetMachineId()

	cmd := exec.Command(opts.Path, "delete", machineId)
	cmd.Env = os.Environ()
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func Init() error {
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	// execute 'multipass version' command
	// as a way to check if multipass is available
	cmd := exec.Command(opts.Path, "version")
	cmd.Env = os.Environ()
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func Start() error {
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machineId := opts.GetMachineId()

	cmd := exec.Command(opts.Path, "start", machineId)
	cmd.Env = os.Environ()
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func Status() error {
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machineId := opts.GetMachineId()

	cmd := exec.Command(opts.Path, "info", machineId)
	cmd.Env = os.Environ()
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func Stop() error {
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machineId := opts.GetMachineId()

	cmd := exec.Command(opts.Path, "stop", machineId)
	cmd.Env = os.Environ()
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
