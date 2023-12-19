package multipass

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/minhio/devpod-provider-multipass/pkg/options"
)

var statusMap = map[string]string{
	"Running":          "Running",
	"Stopped":          "Stopped",
	"Deleted":          "NotFound",
	"Starting":         "Busy",
	"Restarting":       "Busy",
	"Delayed Shutdown": "Busy",
	"Suspending":       "Busy",
	"Suspended":        "Stopped",
	"Unknown":          "NotFound",
}

func Command() error {
	devPodCommand := os.Getenv("COMMAND")
	if devPodCommand == "" {
		return fmt.Errorf("command environment variable is missing")
	}

	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	// example: multipass exec devpod-abc123 -- echo "hello world"
	cmd := exec.Command(opts.Path, "exec", machine.ID, "--", devPodCommand)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

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

	machine := provider.FromEnvironment()

	// example: multipass launch --cpus 2 --disk 40G --memory 2G --name devpod-abc123 lts
	cmd := exec.Command(opts.Path, "launch",
		"--cpus", strconv.Itoa(opts.Cpus),
		"--disk", opts.DiskSize,
		"--memory", opts.Memory,
		"--name", machine.ID,
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

	machine := provider.FromEnvironment()

	// example: multipass delete devpod-abc123
	cmd := exec.Command(opts.Path, "delete", machine.ID)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

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
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

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

	machine := provider.FromEnvironment()

	// example: multipass start devpod-abc123
	cmd := exec.Command(opts.Path, "start", machine.ID)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

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

	machine := provider.FromEnvironment()

	// example: multipass info --format json devpod-abc123
	cmd := exec.Command(opts.Path, "info", "--format", "json", machine.ID)
	cmd.Env = os.Environ()

	output, err := cmd.Output()
	if err != nil {
		return err
	}

	var infoObj map[string]interface{}
	err = json.Unmarshal(output, &infoObj)
	if err != nil {
		return err
	}

	info := infoObj["info"].(map[string]interface{})
	primary := info["primary"].(map[string]interface{})
	state := primary["state"].(string)

	devPodStatus := statusMap[state]
	fmt.Fprint(os.Stdout, devPodStatus)

	return nil
}

func Stop() error {
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	// example: multipass stop devpod-abc123
	cmd := exec.Command(opts.Path, "stop", machine.ID)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
