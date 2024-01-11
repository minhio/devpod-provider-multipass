package multipass

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	CPUS   = "cpus"
	DISK   = "disk"
	MEMORY = "memory"
)

func (c *client) Get(instanceName string, keyName string) (string, error) {
	key := fmt.Sprintf("local.%s.%s", instanceName, keyName)
	args := []string{"get", key}

	log.Default().Printf("get args: %s", args)

	cmd := exec.Command(c.executablePath, args...)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s %s", string(out), err.Error())
	}

	log.Default().Printf("get result: %s", out)

	return string(out), nil
}
