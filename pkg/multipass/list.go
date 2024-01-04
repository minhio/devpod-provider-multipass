package multipass

import (
	"encoding/json"
	"errors"
	"os/exec"
)

type listResult struct {
	List []struct {
		Ipv4    []string `json:"ipv4"`
		Name    string   `json:"name"`
		Release string   `json:"release"`
		State   string   `json:"state"`
	} `json:"list"`
}

func (c *client) List() (*listResult, error) {
	args := []string{"list", "--format", "json"}

	cmd := exec.Command(c.executablePath, args...)
	cmd.Env = c.environ
	cmd.Stdin = c.stdin
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(out) + "\n" + err.Error())
	}

	var result listResult
	err = json.Unmarshal(out, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
