package main

import (
	"os"

	"github.com/hakkdevops/takwin/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
