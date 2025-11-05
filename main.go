package main

import (
	"os"

	"github.com/hakkim/takwin-go/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
