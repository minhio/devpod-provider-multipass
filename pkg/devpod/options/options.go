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
)

type Options struct {
	Path     string // Path to multipass binary
	Image    string // --image arg passed into multipass launch command
	Cpus     int    // --cpus arg passed into multipass launch command
	DiskSize string // --disk arg passed into multipass launch command
	Memory   string // --memory arg passed into multipass launch command
}

func FromEnv() (*Options, error) {
	multipassOptions := &Options{}

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

	multipassOptions.DiskSize, err = fromEnvOrError(MULTIPASS_DISK_SIZE)
	if err != nil {
		return nil, err
	}

	multipassOptions.Memory, err = fromEnvOrError(MULTIPASS_MEMORY)
	if err != nil {
		return nil, err
	}

	return multipassOptions, nil
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
