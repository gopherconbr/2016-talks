package main

import (
	"fmt"
	"os"
	"os/exec"
)

func uptime() error {
	cmd := exec.Command("uptime")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run() // HL1
}

// OMIT

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(2)
}

func main() {
	if err := uptime(); err != nil {
		fatal(err)
	}
}
