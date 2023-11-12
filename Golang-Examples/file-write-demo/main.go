package main

import (
	"fmt"
	"os"
)

func writeFile(name string, text string) error {
	err := os.WriteFile(name, []byte(text), 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	name := "example.txt"
	text := "Oh, hi Mark!"

	err := writeFile(name, text)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("File written successfully.")
}
