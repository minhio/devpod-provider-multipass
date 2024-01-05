package multipass

import (
	"fmt"
	"os"
	"os/exec"
)

func (c *client) Delete(name string) error {
	args := []string{"delete", "--purge", name}

	cmd := exec.Command(c.executablePath, args...)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s %s", string(out), err.Error())
	}

	return nil
}
