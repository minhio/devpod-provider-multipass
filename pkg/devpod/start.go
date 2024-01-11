package devpod

import (
	"fmt"

	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

// devpod calls this to start the multipass instance
func Start() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()
	client, err := multipass.NewClient(opts.Path)
	if err != nil {
		return err
	}

	err = client.Set(machine.ID, multipass.CPUS, fmt.Sprint(opts.Cpus))
	if err != nil {
		return err
	}
	err = client.Set(machine.ID, multipass.DISK, opts.DiskSize)
	if err != nil {
		return err
	}
	err = client.Set(machine.ID, multipass.MEMORY, opts.Memory)
	if err != nil {
		return err
	}

	return client.Start(machine.ID)
}
