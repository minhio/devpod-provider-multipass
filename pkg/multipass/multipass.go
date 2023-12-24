package multipass

import (
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strconv"
)

const (
	RUNNING          = "Running"
	STOPPED          = "Stopped"
	DELETED          = "Deleted"
	STARTING         = "Starting"
	RESTARTING       = "Restarting"
	DELAYED_SHUTDOWN = "Delayed Shutdown"
	SUSPENDING       = "Suspending"
	SUSPENDED        = "Suspended"
	UNKNOWN          = "Unknown"
)

type client struct {
	executablePath string
	env            []string
	stdin          io.Reader
	stdout         io.Writer
	stderr         io.Writer
}

type instance struct {
	Name  string
	State string
	Ipv4  []string
}

type listResult struct {
	List []struct {
		Ipv4    []string `json:"ipv4"`
		Name    string   `json:"name"`
		Release string   `json:"release"`
		State   string   `json:"state"`
	} `json:"list"`
}

type InstanceNotFound struct {
	name string
}

func (i *InstanceNotFound) Error() string {
	return fmt.Sprintf("instance not found: %s", i.name)
}

func NewClient(executablePath string, optsSetters ...OptionSetter) *client {
	opts := &Options{}

	for _, setter := range optsSetters {
		setter(opts)
	}

	return &client{
		executablePath: executablePath,
		env:            opts.Env,
		stdin:          opts.Stdin,
		stdout:         opts.Stdout,
		stderr:         opts.Stderr,
	}
}

func (c *client) List() ([]*instance, error) {
	cmd := exec.Command(c.executablePath,
		"list",
		"--format", "json",
	)
	cmd.Env = c.env
	cmd.Stdin = c.stdin
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var result listResult
	err = json.Unmarshal(out, &result)
	if err != nil {
		return nil, err
	}

	instances := make([]*instance, 0)
	for _, item := range result.List {
		inst := &instance{
			Name:  item.Name,
			State: item.State,
			Ipv4:  item.Ipv4,
		}
		instances = append(instances, inst)
	}

	return instances, nil
}

func (c *client) Launch(name string, cpus int, disk string,
	memory string, cloudInit string, image string) error {

	cmd := exec.Command(c.executablePath,
		"launch",
		"--name", name,
		"--cpus", strconv.Itoa(cpus),
		"--disk", disk,
		"--memory", memory,
		"--cloud-init", cloudInit,
		image,
	)
	cmd.Env = c.env
	cmd.Stdin = c.stdin
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	return cmd.Run()
}

func (c *client) Start(name string) error {
	cmd := exec.Command(c.executablePath,
		"start", name,
	)
	cmd.Env = c.env
	cmd.Stdin = c.stdin
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	return cmd.Run()
}

func (c *client) Stop(name string) error {
	cmd := exec.Command(c.executablePath,
		"stop", name,
	)
	cmd.Env = c.env
	cmd.Stdin = c.stdin
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	return cmd.Run()
}

func (c *client) Delete(name string) error {
	cmd := exec.Command(c.executablePath,
		"delete", "--purge", name,
	)
	cmd.Env = c.env
	cmd.Stdin = c.stdin
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	return cmd.Run()
}

func (c *client) Version() error {
	cmd := exec.Command(c.executablePath,
		"version",
	)
	cmd.Env = c.env
	cmd.Stdin = c.stdin
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	return cmd.Run()
}

func (c *client) GetInstance(name string) (*instance, error) {
	instances, err := c.List()
	if err != nil {
		return nil, err
	}

	for _, inst := range instances {
		if inst.Name == name {
			return inst, nil
		}
	}

	return nil, &InstanceNotFound{name: name}
}
