package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/NeowayLabs/nash"
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s", err)
	os.Exit(1)
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	shell, err := nash.New()

	if err != nil {
		fatal(err)
	}

	for {
		content, err := buf.ReadBytes('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			fatal(err)
		}

		err = shell.Exec("<stdin>", string(content))

		if err != nil {
			fatal(err)
		}
	}
}
