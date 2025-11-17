package main

import (
	"fmt"
	"os"

	"github.com/skushagra/go-gist/internal/cli"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)

	}

	cmd := os.Args[1]

	switch cmd {

	case "push":
		if err := cli.Push(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "fetch":
		if err := cli.Fetch(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`Usage:
  go-gist push <path> <duration>
  go-gist fetch <code>

Examples:
  go-gist push ./ 5h
  go-gist fetch simple-fish-69`)
}