package main

import (
	"fmt"
	"os"
)

func readFile(fileName string) (string, error) {
	bytes, err := os.ReadFile(fileName)
	switch err {
	case nil:
		return string(bytes), nil
	default:
		return "", err
	}
}

func main() {
	content, err := readFile("example.txt")
	switch err {
	case nil:
		fmt.Println(content)
	default:
		fmt.Println("File read failed.")
	}
}
