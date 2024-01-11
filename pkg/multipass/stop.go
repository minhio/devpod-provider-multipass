package multipass

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func (c *client) Stop(name string) error {
	args := []string{"stop", name}

	log.Default().Printf("[multipass] args: %s", args)

	cmd := exec.Command(c.executablePath, args...)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s %s", string(out), err.Error())
	}

	return nil
}
