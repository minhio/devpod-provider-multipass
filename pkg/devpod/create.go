package devpod

import (
	"encoding/base64"
	"fmt"
	"log"
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
  shell: /bin/bash
  groups:
    - docker
    - sudo
  ssh_authorized_keys:
    - ` + pubKey
}

// devpod calls this to create the multipass instance
func Create() error {
	log.Default().Printf("[devpod] create")

	// build the machine context from env vars
	machine := provider.FromEnvironment()

	// get multipass options from env vars
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	// init multipass client
	client, err := multipass.NewClient(opts.Path)
	if err != nil {
		return err
	}

	instanceExists, err := client.IsInstanceExist(machine.ID)
	if err != nil {
		return err
	}

	// only create instance if it doesn't exist
	if instanceExists {
		return fmt.Errorf("instance %s already exist", machine.ID)
	}

	// parse mount args
	mounts, err := parseMountArgs(opts.Mounts)
	if err != nil {
		return err
	}

	// create the machine folder where we will write the cloud-init.yaml
	// and private/public key to
	// this is located in ~/.devpod/contexts/default/machines/<machine id>
	err = os.MkdirAll(machine.Folder, 0755)
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
		multipass.SetMounts(mounts),
		multipass.SetLaunchImage(opts.Image),
	)
	if err != nil {
		return err
	}

	return nil
}
