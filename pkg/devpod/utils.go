package devpod

import (
	"fmt"
	"os"
	"strings"

	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

func parseMountArgs(mountOpt string) []multipass.MountArg {
	mountArgs := make([]multipass.MountArg, 0)

	mounts := strings.Split(mountOpt, ",")
	for _, item := range mounts {
		sourceAndTarget := strings.Split(item, ":")
		if len(sourceAndTarget) == 2 {
			source := sourceAndTarget[0]
			target := sourceAndTarget[1]
			if !strings.HasPrefix(target, "/") {
				target = "/home/devpod/" + target
			}
			mountArgs = append(mountArgs, multipass.MountArg{
				Source: source,
				Target: target,
			})
		}
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
