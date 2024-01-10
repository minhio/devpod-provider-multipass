package devpod

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"

	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/loft-sh/devpod/pkg/ssh"
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

// generate cloud-ini that will be used to launch the multipass instance,
// multipass default user is 'ubuntu', we are leaving that as is,
// but we are adding a 'devpod' user and its public key
func generateCloudInit(pubKey string) string {
	return `
users:
- default
- name: devpod
  sudo: ALL=(ALL) NOPASSWD:ALL
  ssh_authorized_keys:
    - ` + pubKey
}

// devpod calls this to create the multipass instance
func Create() error {
	// get multipass options from env vars
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	// parse mount args
	mounts := parseMountArgs(opts.Mounts)

	// validate mount args
	err = validateMountArgs(mounts...)
	if err != nil {
		return err
	}

	// build the machine context from env vars
	machine := provider.FromEnvironment()

	// create the machine folder where we will write the cloud-init.yaml
	// and private/public key to
	// this is located in ~/.devpod/contexts/default/machines/<machine id>
	err = os.MkdirAll(machine.Folder, 0755)
	if err != nil {
		return err
	}

	// init multipass client
	client, err := multipass.NewClient(opts.Path)
	if err != nil {
		return err
	}

	// generate and write public key (and private key) to machine folder
	publicKeyBase, err := ssh.GetPublicKeyBase(machine.Folder)
	if err != nil {
		return err
	}

	// get public key
	publicKey, err := base64.StdEncoding.DecodeString(publicKeyBase)
	if err != nil {
		return err
	}

	// generate cloud-init with public key
	cloudInitStr := generateCloudInit(string(publicKey))

	// write cloud-init.yaml to machine folder
	cloudInitFilePath := filepath.Join(machine.Folder, "cloud-init.yaml")
	err = os.WriteFile(cloudInitFilePath, []byte(cloudInitStr), 0644)
	if err != nil {
		return err
	}

	// launch the multipass instance
	err = client.Launch(
		multipass.SetLaunchName(machine.ID),
		multipass.SetLaunchCpus(opts.Cpus),
		multipass.SetLaunchDisk(opts.DiskSize),
		multipass.SetLaunchMemory(opts.Memory),
		multipass.SetLaunchCloudInit(cloudInitFilePath),
		multipass.SetLaunchImage(opts.Image),
	)
	if err != nil {
		return err
	}

	// mounts
	mountErr := client.Mount(machine.ID, mounts...)
	if err != nil {
		delErr := client.Delete(machine.ID)
		if delErr != nil {
			return errors.Join(mountErr, delErr)
		}
		return err
	}

	return nil
}
