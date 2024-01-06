package devpod

import (
	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

func Delete() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()
	client, err := multipass.NewClient(opts.Path)
	if err != nil {
		return err
	}

	return client.Delete(machine.ID)
}
