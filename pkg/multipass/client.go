package multipass

import (
	"fmt"
	"strings"
)

type client struct {
	executablePath string
}

func NewClient(executablePath string) *client {
	return &client{
		executablePath: executablePath,
	}
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
