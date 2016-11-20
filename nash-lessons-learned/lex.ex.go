package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/NeowayLabs/nash/scanner"
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func scan(content string) {
	l := scanner.Lex("example", content)

	for tok := range l.Tokens {
		fmt.Println(tok)
	}
}

// OMIT

func main() {
	buf := bufio.NewReader(os.Stdin)

	for {
		content, err := buf.ReadBytes('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			fatal(err)
		}

		scan(string(content))
	}
}
