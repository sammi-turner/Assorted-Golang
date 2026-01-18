package main

import (
	"fmt"
	"os"
)

func writeFile(name string, text string) error {
	err := os.WriteFile(name, []byte(text), 0644)
	switch err {
	case nil:
		return nil
	default:
		return err
	}
}

func main() {
	name := "example.txt"
	text := "Oh, hi Mark!"
	err := writeFile(name, text)
	switch err {
	case nil:
		fmt.Println("File written successfully.")
	default:
		fmt.Println("Error: file write failed.")
	}
}
