package multipass

import "fmt"

const (
	RUNNING          = "Running"
	STOPPED          = "Stopped"
	DELETED          = "Deleted"
	STARTING         = "Starting"
	RESTARTING       = "Restarting"
	DELAYED_SHUTDOWN = "Delayed Shutdown"
	SUSPENDING       = "Suspending"
	SUSPENDED        = "Suspended"
	UNKNOWN          = "Unknown"
)

type instance struct {
	Ipv4  []string
	Name  string
	State string
}

type instanceNotFound struct {
	name string
}

func (i *instanceNotFound) Error() string {
	return fmt.Sprintf("instance not found: %s", i.name)
}
