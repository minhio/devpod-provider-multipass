package multipass

import (
	"fmt"
)

type InstanceNotFound struct {
	name string
}

func (i *InstanceNotFound) Error() string {
	return fmt.Sprintf("instance not found: %s", i.name)
}

func IsInstanceNotFound(err error) bool {
	if _, ok := err.(*InstanceNotFound); ok {
		return true
	}
	return false
}
