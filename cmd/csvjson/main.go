package main

import (
	"os"

	"github.com/pastdev/csvjson/cmd/csvjson/root"
)

func main() {
	if err := root.New().Execute(); err != nil {
		os.Exit(1)
	}
}
