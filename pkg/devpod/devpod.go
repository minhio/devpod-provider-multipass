package devpod

import (
	"fmt"
	"os"

	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

const (
	RUNNING  = "Running"
	STOPPED  = "Stopped"
	NOTFOUND = "NotFound"
	BUSY     = "Busy"
)

var statusMap = map[string]string{
	multipass.RUNNING:          RUNNING,
	multipass.STOPPED:          STOPPED,
	multipass.DELETED:          NOTFOUND,
	multipass.STARTING:         BUSY,
	multipass.RESTARTING:       BUSY,
	multipass.DELAYED_SHUTDOWN: BUSY,
	multipass.SUSPENDING:       BUSY,
	multipass.SUSPENDED:        STOPPED,
	multipass.UNKNOWN:          NOTFOUND,
}

func Command() error {
	devPodCommand := os.Getenv("COMMAND")
	if devPodCommand == "" {
		return fmt.Errorf("command environment variable is missing")
	}

	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	m := multipass.NewMultipass(
		opts.Path,
		multipass.Env(os.Environ()),
		multipass.Stdin(os.Stdin),
		multipass.Stdout(os.Stdout),
		multipass.Stderr(os.Stderr),
	)

	return m.Exec(machine.ID, devPodCommand)
}

func Create() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	m := multipass.NewMultipass(
		opts.Path,
		multipass.Env(os.Environ()),
	)

	return m.Launch(machine.ID, opts.Cpus, opts.DiskSize, opts.Memory, opts.Image)
}

func Delete() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	m := multipass.NewMultipass(
		opts.Path,
		multipass.Env(os.Environ()),
	)

	return m.Delete(machine.ID)
}

func Init() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	m := multipass.NewMultipass(
		opts.Path,
		multipass.Env(os.Environ()),
	)

	return m.Version()
}

func Start() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	m := multipass.NewMultipass(
		opts.Path,
		multipass.Env(os.Environ()),
	)

	return m.Start(machine.ID)
}

func Status() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	m := multipass.NewMultipass(
		opts.Path,
		multipass.Env(os.Environ()),
	)

	instances, err := m.List()
	if err != nil {
		return err
	}

	for _, inst := range instances {
		if inst.Name == machine.ID {
			status := statusMap[inst.State]
			_, err = fmt.Fprint(os.Stdout, status)
			return err
		}
	}

	_, err = fmt.Fprint(os.Stdout, NOTFOUND)
	return err
}

func Stop() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	m := multipass.NewMultipass(
		opts.Path,
		multipass.Env(os.Environ()),
	)

	return m.Stop(machine.ID)
}
