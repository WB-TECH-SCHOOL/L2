package telenet

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var timeout time.Duration

var telnetCmd = &cobra.Command{
	Use:   "telnet [host] [port]",
	Short: "Connect to a host using Telnet protocol",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		port := args[1]
		runTelnet(host, port, timeout)
	},
}

func init() {
	telnetCmd.Flags().DurationVarP(&timeout, "timeout", "t", 10*time.Second, "Connection timeout")
}

func runTelnet(host, port string, timeout time.Duration) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
		os.Exit(2)
	}
	defer conn.Close()

	handleSignals(conn)

	go readFromConnection(conn)
	writeToConnection(conn)
}

func handleSignals(conn net.Conn) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nReceived interrupt signal, closing connection...")
		conn.Close()
		os.Exit(0)
	}()
}

func readFromConnection(conn net.Conn) {
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from connection: %v\n", err)
		os.Exit(2)
	}
}

func writeToConnection(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("EOF received, closing connection...")
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
			os.Exit(2)
		}
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to connection: %v\n", err)
			os.Exit(2)
		}
	}
}

func Execute() error {
	return telnetCmd.Execute()
}
