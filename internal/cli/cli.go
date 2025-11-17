package cli

import (
	"errors"
	"fmt"
)

func Push(args []string) error {
	if len(args) < 2 {
		return errors.New("push required a file path <path> and a duration <duration>. example: go-gist push ./ 5h")
	}

	path := args[0]
	duration := args[1]

	// TODO: Implement the logic to push the gist

	fmt.Printf("Pushed gist from path: %s with duration: %s\n", path, duration)
	return nil

}

func Fetch(args []string) error {
	if len(args) < 1 {
		return errors.New("fetch required a gist code <code>. example: go-gist fetch simple-fish-69")
	}

	code := args[0]

	// TODO: Implement the logic to fetch the gist

	fmt.Printf("Fetched gist with code: %s\n", code)
	return nil

}