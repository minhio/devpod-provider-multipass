package multipass

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func (c *client) Set(instanceName string, keyName string, value string) error {
	key := fmt.Sprintf("local.%s.%s", instanceName, keyName)
	args := []string{"set", key + "=" + value}

	log.Default().Printf("[multipass] args: %s", args)

	cmd := exec.Command(c.executablePath, args...)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s %s", string(out), err.Error())
	}

	return nil
}
