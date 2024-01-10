package multipass

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type client struct {
	executablePath string
}

func NewClient(executablePath string) (*client, error) {
	log.Default().Printf("executable path: %s", executablePath)

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
	log.Default().Printf("get instance: %s", name)

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
