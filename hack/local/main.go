package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("./release/provider.yaml")
	if err != nil {
		panic(err)
	}

	strToReplace := fmt.Sprintf("https://github.com/minhio/devpod-provider-multipass/releases/download/%s", os.Args[1])
	replaced := strings.Replace(string(content), strToReplace, os.Args[2], -1)
	fmt.Print(replaced)
}
