package multipass

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type launchArgs struct {
	name      string
	cpus      int
	disk      string
	memory    string
	cloudInit string
	image     string
}

type argSetter func(*launchArgs)

func (c *client) Launch(argSetters ...argSetter) error {
	launchArgz := &launchArgs{}

	for _, setter := range argSetters {
		setter(launchArgz)
	}

	args := []string{"launch"}

	if launchArgz.name != "" {
		args = append(args, "--name", launchArgz.name)
	}

	if launchArgz.cpus != 0 {
		args = append(args, "--cpus", strconv.Itoa(launchArgz.cpus))
	}

	if launchArgz.disk != "" {
		args = append(args, "--disk", launchArgz.disk)
	}

	if launchArgz.memory != "" {
		args = append(args, "--memory", launchArgz.memory)
	}

	if launchArgz.cloudInit != "" {
		args = append(args, "--cloud-init", launchArgz.cloudInit)
	}

	if launchArgz.image != "" {
		args = append(args, launchArgz.image)
	}

	cmd := exec.Command(c.executablePath, args...)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s %s", string(out), err.Error())
	}

	return nil
}

func SetLaunchName(name string) argSetter {
	return func(args *launchArgs) {
		args.name = name
	}
}

func SetLaunchCpus(cpus int) argSetter {
	return func(args *launchArgs) {
		args.cpus = cpus
	}
}

func SetLaunchDisk(disk string) argSetter {
	return func(args *launchArgs) {
		args.disk = disk
	}
}

func SetLaunchMemory(memory string) argSetter {
	return func(args *launchArgs) {
		args.memory = memory
	}
}

func SetLaunchCloudInit(cloudInit string) argSetter {
	return func(args *launchArgs) {
		args.cloudInit = cloudInit
	}
}

func SetLaunchImage(image string) argSetter {
	return func(args *launchArgs) {
		args.image = image
	}
}
