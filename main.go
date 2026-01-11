package main

import (
	"os"

	"github.com/xrendan/ai-starter/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
