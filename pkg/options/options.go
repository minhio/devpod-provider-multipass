package options

import (
	"fmt"
	"os"
	"strconv"
)

const (
	MULTIPASS_PATH      = "MULTIPASS_PATH"
	MULTIPASS_IMAGE     = "MULTIPASS_IMAGE"
	MULTIPASS_CPUS      = "MULTIPASS_CPUS"
	MULTIPASS_DISK_SIZE = "MULTIPASS_DISK_SIZE"
	MULTIPASS_MEMORY    = "MULTIPASS_MEMORY"

	DEVPOD           = "DEVPOD"
	DEVPOD_OS        = "DEVPOD_OS"
	DEVPOD_ARCH      = "DEVPOD_ARCH"
	MACHINE_ID       = "MACHINE_ID"
	MACHINE_FOLDER   = "MACHINE_FOLDER"
	MACHINE_CONTEXT  = "MACHINE_CONTEXT"
	MACHINE_PROVIDER = "MACHINE_PROVIDER"
)

type MultipassOptions struct {
	Path     string
	Image    string
	Cpus     int
	DiskSize int
	Memory   int
}

type DevPodOptions struct {
	DevPodCliPath   string
	DevPodOs        string
	DevPodArch      string
	MachineId       string
	MachineFolder   string
	MachineContext  string
	MachineProvider string
}

func MultipassOptionsFromEnv() (*MultipassOptions, error) {
	multipassOptions := &MultipassOptions{}

	var err error

	multipassOptions.Path, err = fromEnvOrError(MULTIPASS_PATH)
	if err != nil {
		return nil, err
	}

	multipassOptions.Image, err = fromEnvOrError(MULTIPASS_IMAGE)
	if err != nil {
		return nil, err
	}

	cpus, err := fromEnvOrError(MULTIPASS_CPUS)
	if err != nil {
		return nil, err
	}

	multipassOptions.Cpus, err = strconv.Atoi(cpus)
	if err != nil {
		return nil, err
	}

	diskSize, err := fromEnvOrError(MULTIPASS_DISK_SIZE)
	if err != nil {
		return nil, err
	}

	multipassOptions.DiskSize, err = strconv.Atoi(diskSize)
	if err != nil {
		return nil, err
	}

	memory, err := fromEnvOrError(MULTIPASS_MEMORY)
	if err != nil {
		return nil, err
	}

	multipassOptions.Memory, err = strconv.Atoi(memory)
	if err != nil {
		return nil, err
	}

	return multipassOptions, nil
}

func DevPodOptionsFromEnv() (*DevPodOptions, error) {
	devPodOptions := &DevPodOptions{}

	var err error

	devPodOptions.DevPodCliPath, err = fromEnvOrError(DEVPOD)
	if err != nil {
		return nil, err
	}

	devPodOptions.DevPodOs, err = fromEnvOrError(DEVPOD_OS)
	if err != nil {
		return nil, err
	}

	devPodOptions.DevPodArch, err = fromEnvOrError(DEVPOD_ARCH)
	if err != nil {
		return nil, err
	}

	devPodOptions.MachineId, err = fromEnvOrError(MACHINE_ID)
	if err != nil {
		return nil, err
	}

	devPodOptions.MachineFolder, err = fromEnvOrError(MACHINE_FOLDER)
	if err != nil {
		return nil, err
	}

	devPodOptions.MachineContext, err = fromEnvOrError(MACHINE_CONTEXT)
	if err != nil {
		return nil, err
	}

	devPodOptions.MachineProvider, err = fromEnvOrError(MACHINE_PROVIDER)
	if err != nil {
		return nil, err
	}

	return devPodOptions, nil
}

func fromEnvOrError(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return "", fmt.Errorf(
			"couldn't find option %s in environment, please make sure %s is defined",
			name,
			name,
		)
	}

	return val, nil
}
