package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/NeowayLabs/nash/parser"
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func parse(content string) {
	parser := parser.NewParser("example", content)
	tree, err := parser.Parse()

	if err != nil {
		fatal(err)
	}

	for _, node := range tree.Root.Nodes {
		fmt.Printf("Node type: %s, toString: %s\n", node.Type(), node.String())
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

		parse(string(content))
	}
}
