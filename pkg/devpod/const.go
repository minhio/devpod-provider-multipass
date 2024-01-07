package devpod

import "github.com/minhio/devpod-provider-multipass/pkg/multipass"

const (
	RUNNING  = "Running"
	STOPPED  = "Stopped"
	NOTFOUND = "NotFound"
	BUSY     = "Busy"
)

// devpod expects four status above, multipass has nine
// here we are mapping multipass instance status to what devpod expects
var statusMap = map[string]string{
	multipass.STATE_RUNNING:          RUNNING,
	multipass.STATE_STOPPED:          STOPPED,
	multipass.STATE_DELETED:          NOTFOUND,
	multipass.STATE_STARTING:         BUSY,
	multipass.STATE_RESTARTING:       BUSY,
	multipass.STATE_DELAYED_SHUTDOWN: BUSY,
	multipass.STATE_SUSPENDING:       BUSY,
	multipass.STATE_SUSPENDED:        STOPPED,
	multipass.STATE_UNKNOWN:          NOTFOUND,
}
