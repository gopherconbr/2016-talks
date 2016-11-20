package main

import (
	"fmt"
	"os"
	"syscall"
)

func uptime() error {
	pid, err := syscall.ForkExec("/usr/bin/uptime", []string{"uptime"}, &syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()}, // HL1
	})

	if err != nil {
		return err
	}

	var wstatus syscall.WaitStatus

	// maybe handle EINTR
	if _, err = syscall.Wait4(pid, &wstatus, 0, nil); err != nil { // HL2
		return err
	}

	if wstatus.ExitStatus() != 0 { // HL3
		return fmt.Errorf("command failed with status %d", wstatus.ExitStatus())
	}

	return nil
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
