package devpod

import (
	"fmt"
	"os"

	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

// devpod calls this to get the status of the multipass instance
// and expect the response to be returned via stdout
func Status() error {
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
