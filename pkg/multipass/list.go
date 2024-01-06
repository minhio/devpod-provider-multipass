package multipass

import (
	"encoding/json"
	"fmt"
	"os"
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
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s %s", string(out), err.Error())
	}

	var result listResult
	err = json.Unmarshal(out, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
