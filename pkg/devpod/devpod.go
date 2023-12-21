package devpod

import (
	"fmt"
	"os"

	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/minhio/devpod-provider-multipass/pkg/devpod/options"
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

	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	devPodCommand = fmt.Sprintf("sh -c \"%s\"", devPodCommand)
	return multipass.NewMultipass(opts.Path).Exec(machine.ID, devPodCommand)
}

func Create() error {
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	return multipass.NewMultipass(opts.Path).Launch(machine.ID, opts.Cpus,
		opts.DiskSize, opts.Memory, opts.Image)
}

func Delete() error {
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	return multipass.NewMultipass(opts.Path).Delete(machine.ID)
}

func Init() error {
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	return multipass.NewMultipass(opts.Path).Version()
}

func Start() error {
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	return multipass.NewMultipass(opts.Path).Start(machine.ID)
}

func Status() error {
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	instances, err := multipass.NewMultipass(opts.Path).List()
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
	opts, err := options.FromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	return multipass.NewMultipass(opts.Path).Stop(machine.ID)
}
