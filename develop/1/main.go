package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	exactTime, err := ntp.Time("time.google.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching NTP time: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Current system time:", time.Now().Format(time.RFC1123))

	fmt.Println("Exact NTP time:", exactTime.Format(time.RFC1123))
}
