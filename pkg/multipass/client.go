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
	log.Default().Printf("[multipass] path: %s", executablePath)

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
	log.Default().Printf("[multipass] getting instance: %s", name)

	infoResult, err := c.Info(name)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("instance \"%s\" does not exist", name)) {
			log.Default().Printf("[multipass] instance not found: %s", name)
			return nil, &InstanceNotFound{name: name}
		}
		return nil, err
	}

	instInfo := infoResult.Info[name]

	return &instInfo, nil
}

func (c *client) IsInstanceExist(name string) (bool, error) {
	_, err := c.GetInstance(name)
	if err != nil {
		if IsInstanceNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
