package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	args := os.Args[1:]

	switch len(args) {
	case 0:
		return errors.New("No argument supplied")
	case 1:
		fmt.Println(args[0])
		return nil
	default:
		return errors.New("Only one argument is allowed")
	}
}
