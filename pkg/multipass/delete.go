package multipass

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func (c *client) Delete(name string) error {
	args := []string{"delete", "--purge", name}

	log.Default().Printf("[multipass] %s", args)

	cmd := exec.Command(c.executablePath, args...)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s %s", string(out), err.Error())
	}

	return nil
}
