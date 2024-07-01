package main

import (
	"os"

	"github.com/rarimo/geo-points-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
