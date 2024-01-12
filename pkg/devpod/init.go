package devpod

import (
	"log"

	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

// devpod calls this when the provider is being added,
// here we are just invoking the 'multipass version' command
// as a way to ensure that multipass is reachable
func Init() error {
	log.Default().Printf("[devpod] init")

	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	// parse mount args
	_, err = parseMountArgs(opts.Mounts)
	if err != nil {
		return err
	}

	client, err := multipass.NewClient(opts.Path)
	if err != nil {
		return err
	}

	_, err = client.Version()
	if err != nil {
		return err
	}

	return nil
}
