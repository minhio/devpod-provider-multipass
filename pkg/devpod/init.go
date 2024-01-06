package devpod

import (
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

func Init() error {
	opts, err := OptsFromEnv()
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
