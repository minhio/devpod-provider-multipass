package devpod

import (
	"fmt"
	"os"
	"strings"

	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

func parseMountArgs(mountOpt string) []multipass.MountArg {
	mountArgs := make([]multipass.MountArg, 0)

	if mountOpt == "" {
		return mountArgs
	}

	mountsFromOpt := strings.Split(mountOpt, ",")

	for _, mount := range mountsFromOpt {
		sourceAndTarget := strings.Split(mount, ":")

		source := sourceAndTarget[0]
		var target string

		if len(sourceAndTarget) == 1 {
			target = source
		} else if len(sourceAndTarget) == 2 {
			target = sourceAndTarget[1]
		}

		if !strings.HasPrefix(target, "/") {
			target = "/home/devpod/" + target
		}

		mountArgs = append(mountArgs, multipass.MountArg{
			Source: source,
			Target: target,
		})
	}

	return mountArgs
}

func validateMountArgs(mountArgs ...multipass.MountArg) error {
	for _, arg := range mountArgs {
		if _, err := os.Stat(arg.Source); os.IsNotExist(err) {
			return fmt.Errorf("%s does not exist", arg.Source)
		}
		if !strings.HasPrefix(arg.Target, "/") {
			return fmt.Errorf("%s is not absolute path", arg.Target)
		}
	}
	return nil
}
