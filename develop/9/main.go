package main

import (
	"fmt"
	"os"
	"wget/wget"
)

func main() {
	if err := wget.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
