package devpod

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/loft-sh/devpod/pkg/ssh"
	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
	"github.com/pkg/errors"
)

const (
	RUNNING  = "Running"
	STOPPED  = "Stopped"
	NOTFOUND = "NotFound"
	BUSY     = "Busy"
)

var statusMap = map[string]string{
	multipass.RUNNING:          RUNNING,
	multipass.STOPPED:          STOPPED,
	multipass.DELETED:          NOTFOUND,
	multipass.STARTING:         BUSY,
	multipass.RESTARTING:       BUSY,
	multipass.DELAYED_SHUTDOWN: BUSY,
	multipass.SUSPENDING:       BUSY,
	multipass.SUSPENDED:        STOPPED,
	multipass.UNKNOWN:          NOTFOUND,
}

func Command() error {
	devPodCommand := os.Getenv("COMMAND")
	if devPodCommand == "" {
		return fmt.Errorf("command environment variable is missing")
	}

	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	multipassClient := multipass.NewClient(
		opts.Path,
		multipass.Env(os.Environ()),
	)

	instance, err := multipassClient.GetInstance(machine.ID)
	if err != nil {
		return err
	}

	privateKey, err := ssh.GetPrivateKeyRawBase(machine.Folder)
	if err != nil {
		return fmt.Errorf("load private key: %w", err)
	}

	sshClient, err := ssh.NewSSHClient("devpod", instance.Ipv4[0]+":22", privateKey)
	if err != nil {
		return errors.Wrap(err, "create ssh client")
	}
	defer sshClient.Close()

	return ssh.Run(context.Background(), sshClient, devPodCommand, os.Stdin, os.Stdout, os.Stderr)
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

	multipassClient := multipass.NewClient(
		opts.Path,
		multipass.Env(os.Environ()),
	)

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

	return multipassClient.Launch(machine.ID, opts.Cpus, opts.DiskSize,
		opts.Memory, cloudInitFilePath, opts.Image)
}

func Delete() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	multipassClient := multipass.NewClient(
		opts.Path,
		multipass.Env(os.Environ()),
	)

	return multipassClient.Delete(machine.ID)
}

func Init() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	m := multipass.NewClient(
		opts.Path,
		multipass.Env(os.Environ()),
	)

	return m.Version()
}

func Start() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	multipassClient := multipass.NewClient(
		opts.Path,
		multipass.Env(os.Environ()),
	)

	return multipassClient.Start(machine.ID)
}

func Status() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	multipassClient := multipass.NewClient(
		opts.Path,
		multipass.Env(os.Environ()),
	)

	instance, err := multipassClient.GetInstance(machine.ID)
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

func Stop() error {
	opts, err := OptsFromEnv()
	if err != nil {
		return err
	}

	machine := provider.FromEnvironment()

	multipassClient := multipass.NewClient(
		opts.Path,
		multipass.Env(os.Environ()),
	)

	return multipassClient.Stop(machine.ID)
}

func generateCloudInit(pubKey string) string {
	return `
users:
- default
- name: devpod
  sudo: ALL=(ALL) NOPASSWD:ALL
  ssh_authorized_keys:
    - ` + pubKey
}
