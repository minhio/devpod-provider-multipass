package multipass

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type mountInfo struct {
	SourcePath string `json:"source_path"`
}

type instanceInfo struct {
	Name   string
	Ipv4   []string             `json:"ipv4"`
	Mounts map[string]mountInfo `json:"mounts"`
	State  string               `json:"state"`
}

type infoResult struct {
	Info map[string]instanceInfo `json:"info"`
}

func (c *client) Info(name string) (*infoResult, error) {
	args := []string{"info", "--format", "json", name}

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

	if instInfo, ok := result.Info[name]; ok {
		instInfo.Name = name
	}

	return &result, nil
}
