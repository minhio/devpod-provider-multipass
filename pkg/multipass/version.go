package multipass

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type versionResult struct {
	Multipass  string `json:"multipass"`
	Multipassd string `json:"multipassd"`
}

func (c *client) Version() (*versionResult, error) {
	args := []string{"version", "--format", "json"}

	log.Default().Printf("[multipass] args: %s", args)

	cmd := exec.Command(c.executablePath, args...)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s %s", string(out), err.Error())
	}

	var result versionResult
	err = json.Unmarshal(out, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
