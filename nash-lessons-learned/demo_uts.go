// This program is the Go version of namespace example program
// demo_uts.c of Linux Namespace series of articles:
// https://lwn.net/Articles/531381/
//
// Go has serious problems with fork, exec, and clone syscalls because
// of the runtime threads. The idea to implement the same behaviour of the
// C version was using os.Exec'ing the same program passing a special argument
// in the end of argument list. This idea was stolen by R. Minnich
// u-root project.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

// Converts a C string into a Go string
func getstr(cstr [65]int8) string {
	b := make([]byte, 0, 65)

	for _, i := range cstr {
		if i == 0 {
			break
		}

		b = append(b, byte(i))
	}

	return string(b)
}

func getuts() string {
	uts := syscall.Utsname{}

	err := syscall.Uname(&uts)

	if err != nil {
		fatal(err)
	}

	return getstr(uts.Nodename)
}

func fork() {
	cmd := &exec.Cmd{
		Path: os.Args[0],
		Args: append([]string{"_demo_uts_"}, os.Args[1:]...),
	}

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER,
	}

	cmd.SysProcAttr.UidMappings = []syscall.SysProcIDMap{{
		ContainerID: 0,
		HostID:      os.Getuid(),
		Size:        1,
	}}

	cmd.SysProcAttr.GidMappings = []syscall.SysProcIDMap{{
		ContainerID: 0,
		HostID:      os.Getgid(),
		Size:        1,
	}}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fatal(err)
	}

	fmt.Printf("Parent uts.nodename = %s\n", getuts())
}

func updateHostname(hostname string) {
	err := syscall.Sethostname([]byte(hostname))

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Child uts.nodename = %s\n", getuts())
}

func main() {
	if len(os.Args[0]) > 0 && os.Args[0] != "_demo_uts_" {
		// This is the parent code
		fork()
	} else {
		// Child namespace'd code
		updateHostname("Gophercon Brazil 2016")
	}
}
