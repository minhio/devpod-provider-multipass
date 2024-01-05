package multipass

import "fmt"

type InstanceNotFound struct {
	name string
}

func (i *InstanceNotFound) Error() string {
	return fmt.Sprintf("instance not found: %s", i.name)
}
