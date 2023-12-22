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

type multipass struct {
	executablePath string
	env            []string
	stdin          io.Reader
	stdout         io.Writer
	stderr         io.Writer
}

type instance struct {
	Name  string
	State string
}

type listResult struct {
	List []struct {
		Ipv4    []string `json:"ipv4"`
		Name    string   `json:"name"`
		Release string   `json:"release"`
		State   string   `json:"state"`
	} `json:"list"`
}

func NewMultipass(executablePath string, optsSetters ...OptionSetter) *multipass {
	opts := &Options{}

	for _, setter := range optsSetters {
		setter(opts)
	}

	return &multipass{
		executablePath: executablePath,
		env:            opts.Env,
		stdin:          opts.Stdin,
		stdout:         opts.Stdout,
		stderr:         opts.Stderr,
	}
}

func (m multipass) List() ([]instance, error) {
	cmd := exec.Command(m.executablePath,
		"list",
		"--format", "json",
	)
	cmd.Env = m.env
	cmd.Stdin = m.stdin
	cmd.Stdout = m.stdout
	cmd.Stderr = m.stderr

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var result listResult
	err = json.Unmarshal(out, &result)
	if err != nil {
		return nil, err
	}

	instances := make([]instance, 0)
	for _, item := range result.List {
		inst := instance{
			Name:  item.Name,
			State: item.State,
		}
		instances = append(instances, inst)
	}

	return instances, nil
}

func (m multipass) Launch(name string, cpus int, disk string,
	memory string, image string) error {

	cmd := exec.Command(m.executablePath,
		"launch",
		"--name", name,
		"--cpus", strconv.Itoa(cpus),
		"--disk", disk,
		"--memory", memory,
		image,
	)
	cmd.Env = m.env
	cmd.Stdin = m.stdin
	cmd.Stdout = m.stdout
	cmd.Stderr = m.stderr

	return cmd.Run()
}

func (m multipass) Start(name string) error {
	cmd := exec.Command(m.executablePath,
		"start", name,
	)
	cmd.Env = m.env
	cmd.Stdin = m.stdin
	cmd.Stdout = m.stdout
	cmd.Stderr = m.stderr

	return cmd.Run()
}

func (m multipass) Stop(name string) error {
	cmd := exec.Command(m.executablePath,
		"stop", name,
	)
	cmd.Env = m.env
	cmd.Stdin = m.stdin
	cmd.Stdout = m.stdout
	cmd.Stderr = m.stderr

	return cmd.Run()
}

func (m multipass) Delete(name string) error {
	cmd := exec.Command(m.executablePath,
		"delete", "--purge", name,
	)
	cmd.Env = m.env
	cmd.Stdin = m.stdin
	cmd.Stdout = m.stdout
	cmd.Stderr = m.stderr

	return cmd.Run()
}

func (m multipass) Exec(name string, command string) error {
	cmd := exec.Command(m.executablePath,
		"exec",
		name,
		"--",
		fmt.Sprintf("sh -c '%s'", command),
	)
	cmd.Env = m.env
	cmd.Stdin = m.stdin
	cmd.Stdout = m.stdout
	cmd.Stderr = m.stderr

	return cmd.Run()
}

func (m multipass) Version() error {
	cmd := exec.Command(m.executablePath,
		"version",
	)
	cmd.Env = m.env
	cmd.Stdin = m.stdin
	cmd.Stdout = m.stdout
	cmd.Stderr = m.stderr

	return cmd.Run()
}
