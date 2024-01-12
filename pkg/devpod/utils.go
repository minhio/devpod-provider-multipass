package devpod

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/minhio/devpod-provider-multipass/pkg/multipass"
)

func parseMountArgs(mountOpt string) ([]multipass.MountArg, error) {
	if mountOpt == "" {
		return nil, nil
	}

	mountArgs := make([]multipass.MountArg, 0)

	for _, mount := range strings.Split(mountOpt, ",") {
		sourceAndTarget := strings.Split(mount, ":")

		source := filepath.Join(sourceAndTarget[0])
		if _, err := os.Stat(source); os.IsNotExist(err) {
			return nil, fmt.Errorf("%s does not exist", source)
		}

		target := ""
		if len(sourceAndTarget) == 2 {
			if strings.HasPrefix(sourceAndTarget[1], "/") {
				target = filepath.Join(sourceAndTarget[1])
			} else {
				target = filepath.Join("/", "home", "devpod", sourceAndTarget[1])
			}
		}

		mountArgs = append(mountArgs, multipass.MountArg{
			Source: source,
			Target: target,
		})
	}

	return mountArgs, nil
}
