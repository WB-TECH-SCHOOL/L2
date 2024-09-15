package main

import (
	"bufio"
	"fmt"
	"os"
	"shell/shell"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("shell> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "\\quit" {
			break
		}
		shell.Execute(line)
	}
}
