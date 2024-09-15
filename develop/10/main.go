package main

import (
	"fmt"
	"os"
	"telenet/telenet"
)

func main() {
	if err := telenet.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
