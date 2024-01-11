package multipass

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type mountInfo struct {
	SourcePath string `json:"source_path"`
}

type instanceInfo struct {
	Ipv4   []string             `json:"ipv4"`
	Mounts map[string]mountInfo `json:"mounts"`
	State  string               `json:"state"`
}

type infoResult struct {
	Info map[string]instanceInfo `json:"info"`
}

func (c *client) Info(name string) (*infoResult, error) {
	args := []string{"info", "--format", "json", name}

	log.Default().Printf("[multipass] %s", args)

	cmd := exec.Command(c.executablePath, args...)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s %s", string(out), err.Error())
	}

	var result infoResult
	err = json.Unmarshal(out, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
