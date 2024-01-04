package multipass

import (
	"encoding/json"
	"errors"
	"os/exec"
)

type versionResult struct {
	Multipass  string `json:"multipass"`
	Multipassd string `json:"multipassd"`
}

func (c *client) Version() (*versionResult, error) {
	args := []string{"version"}

	cmd := exec.Command(c.executablePath, args...)
	cmd.Env = c.environ
	cmd.Stdin = c.stdin
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(out) + "\n" + err.Error())
	}

	var result versionResult
	err = json.Unmarshal(out, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
