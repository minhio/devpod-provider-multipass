package multipass

import (
	"fmt"
	"os/exec"
	"strings"
)

type client struct {
	executablePath string
}

func NewClient(executablePath string) (*client, error) {
	_, err := exec.LookPath(executablePath)
	if err != nil {
		return nil, err
	}

	client := client{
		executablePath: executablePath,
	}

	return &client, nil
}

func (c *client) GetInstance(name string) (*instanceInfo, error) {
	infoResult, err := c.Info(name)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("instance \"%s\" does not exist", name)) {
			return nil, &InstanceNotFound{name: name}
		}
		return nil, err
	}

	instInfo := infoResult.Info[name]
	return &instInfo, nil
}
