package devpod

import (
	"fmt"
	"log"
	"os"

	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

// devpod calls this to get the status of the multipass instance
// and expect the response to be returned via stdout
func Status() error {
	log.Default().Printf("[devpod] status")

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

	log.Default().Printf("[devpod] status: %s", status)
	_, err = fmt.Fprint(os.Stdout, status)
	return err
}
