package multipass

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type MountArg struct {
	Source string
	Target string
}

func (c *client) Mount(name string, mounts ...MountArg) error {
	for _, mount := range mounts {
		err := c.mount(name, mount)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *client) mount(name string, mount MountArg) error {
	args := []string{"mount", mount.Source, name + ":" + mount.Target}

	log.Default().Printf("[multipass] %s", args)

	cmd := exec.Command(c.executablePath, args...)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s %s", string(out), err.Error())
	}

	return nil
}
