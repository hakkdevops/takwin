package main

import (
	"os"

	"github.com/hakkim/takwin/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
