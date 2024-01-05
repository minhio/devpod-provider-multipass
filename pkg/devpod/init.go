package devpod

import (
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

func Init() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	client := multipass.NewClient(opts.Path)

	_, err = client.Version()
	if err != nil {
		return err
	}

	return nil
}
