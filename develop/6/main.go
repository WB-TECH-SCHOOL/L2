package main

import (
	"cut/cut"
	"fmt"
	"os"
)

func main() {
	if err := cut.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
