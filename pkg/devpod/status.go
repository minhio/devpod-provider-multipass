package devpod

import (
	"fmt"
	"os"

	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

func Status() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()
	client := multipass.NewClient(opts.Path)

	instance, err := client.GetInstance(machine.ID)
	if err != nil {
		if _, ok := err.(*multipass.InstanceNotFound); ok {
			_, err = fmt.Fprint(os.Stdout, NOTFOUND)
			return err
		}
		return err
	}

	status := statusMap[instance.State]
	_, err = fmt.Fprint(os.Stdout, status)
	return err
}
