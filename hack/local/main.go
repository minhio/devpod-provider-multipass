package main

import (
	"fmt"
	"os"
	"strings"
)

// calls from the /hack/build-local.sh script
// this replaces the github path to the binaries in /release/provider.yaml
// with the local path, useful for local development
func main() {
	content, err := os.ReadFile("./release/provider.yaml")
	if err != nil {
		panic(err)
	}

	strToReplace := fmt.Sprintf("https://github.com/minhio/devpod-provider-multipass/releases/download/%s", os.Args[1])
	replaced := strings.Replace(string(content), strToReplace, os.Args[2], -1)
	fmt.Print(replaced)
}
