package multipass

import (
	"fmt"
	"log"
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
	mounts    []MountArg
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

	if len(launchArgz.mounts) > 0 {
		for _, mount := range launchArgz.mounts {
			if mount.Target == "" {
				args = append(args, "--mount", mount.Source)
			} else {
				args = append(args, "--mount", mount.Source+":"+mount.Target)
			}
		}
	}

	if launchArgz.image != "" {
		args = append(args, launchArgz.image)
	}

	log.Default().Printf("[multipass] %s", args)

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

func SetMounts(mounts []MountArg) argSetter {
	return func(args *launchArgs) {
		args.mounts = mounts
	}
}
