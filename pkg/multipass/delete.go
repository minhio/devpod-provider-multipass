package multipass

import (
	"errors"
	"os/exec"
)

func (c *client) Delete(name string) error {
	args := []string{"delete", "--purge", name}

	cmd := exec.Command(c.executablePath, args...)
	cmd.Env = c.environ
	cmd.Stdin = c.stdin
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out) + "\n" + err.Error())
	}

	return nil
}