package devpod

import (
	"fmt"
	"log"

	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

// devpod calls this to start the multipass instance
func Start() error {
	log.Default().Printf("[devpod] start")

	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()
	client, err := multipass.NewClient(opts.Path)
	if err != nil {
		return err
	}

	instance, err := client.GetInstance(machine.ID)
	if err != nil {
		return err
	}

	if instance.State == multipass.STATE_STOPPED {
		err := client.Set(machine.ID, multipass.CPUS, fmt.Sprint(opts.Cpus))
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
	}

	return client.Start(machine.ID)
}
