package main

import (
	"os"

	"github.com/cernbox/ocis-canary/pkg/command"
)

func main() {
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
