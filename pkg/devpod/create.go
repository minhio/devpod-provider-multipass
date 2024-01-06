package devpod

import (
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/loft-sh/devpod/pkg/ssh"
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

func generateCloudInit(pubKey string) string {
	return `
users:
- default
- name: devpod
  sudo: ALL=(ALL) NOPASSWD:ALL
  ssh_authorized_keys:
    - ` + pubKey
}

func Create() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	err = os.MkdirAll(machine.Folder, 0755)
	if err != nil {
		return err
	}

	client, err := multipass.NewClient(opts.Path)
	if err != nil {
		return err
	}

	publicKeyBase, err := ssh.GetPublicKeyBase(machine.Folder)
	if err != nil {
		return err
	}

	publicKey, err := base64.StdEncoding.DecodeString(publicKeyBase)
	if err != nil {
		return err
	}

	cloudInitStr := generateCloudInit(string(publicKey))
	cloudInitFilePath := filepath.Join(machine.Folder, "cloud-init.yaml")

	err = os.WriteFile(cloudInitFilePath, []byte(cloudInitStr), 0644)
	if err != nil {
		return err
	}

	return client.Launch(
		multipass.SetLaunchName(machine.ID),
		multipass.SetLaunchCpus(opts.Cpus),
		multipass.SetLaunchDisk(opts.DiskSize),
		multipass.SetLaunchMemory(opts.Memory),
		multipass.SetLaunchCloudInit(cloudInitFilePath),
		multipass.SetLaunchImage(opts.Image),
	)
}
