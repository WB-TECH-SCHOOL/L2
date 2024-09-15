package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Execute(input string) {
	commands := strings.Split(input, "|")
	var prevStdout io.ReadCloser

	for i, cmdStr := range commands {
		cmdArgs := strings.Fields(strings.TrimSpace(cmdStr))
		if len(cmdArgs) == 0 {
			continue
		}

		switch cmdArgs[0] {
		case "cd":
			changeDirectory(cmdArgs)
			return
		case "pwd":
			printWorkingDirectory()
			return
		case "echo":
			echo(cmdArgs)
			return
		case "kill":
			killProcess(cmdArgs)
			return
		case "ps":
			listProcesses()
			return
		}

		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

		if prevStdout != nil {
			cmd.Stdin = prevStdout
		}

		if i < len(commands)-1 {
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error creating pipe:", err)
				return
			}
			prevStdout = stdout
		} else {
			cmd.Stdout = os.Stdout
		}

		if err := cmd.Start(); err != nil {
			fmt.Fprintln(os.Stderr, "Error starting command:", err)
			return
		}

		if err := cmd.Wait(); err != nil {
			fmt.Fprintln(os.Stderr, "Error waiting for command:", err)
			return
		}
	}
}

func changeDirectory(args []string) {
	if len(args) < 2 {
		fmt.Println("cd: missing argument")
		return
	}
	if err := os.Chdir(args[1]); err != nil {
		fmt.Fprintln(os.Stderr, "cd error:", err)
	}
}

func printWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "pwd error:", err)
		return
	}
	fmt.Println(dir)
}

func echo(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func killProcess(args []string) {
	if len(args) < 2 {
		fmt.Println("kill: missing argument")
		return
	}
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "kill error:", err)
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, "kill error:", err)
		return
	}
	if err := process.Kill(); err != nil {
		fmt.Fprintln(os.Stderr, "kill error:", err)
	}
}

func listProcesses() {
	cmd := exec.Command("tasklist")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "ps error:", err)
	}
}
