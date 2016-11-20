package main

import (
	"fmt"
	"os"
)

func uptime() error {
	procAttr := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, // HL1
	}

	process, err := os.StartProcess("/usr/bin/uptime",
		[]string{"uptime"},
		&procAttr,
	)

	if err != nil {
		return err
	}

	if pstatus, err := process.Wait(); err != nil { // HL2
		return err
	} else if !pstatus.Success() { // HL3
		// Exact status code is system-dependent
		return fmt.Errorf("Failed to execute uptime")
	}

	return nil
}

// OMIT

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func main() {
	if err := uptime(); err != nil {
		fatal(err)
	}
}
