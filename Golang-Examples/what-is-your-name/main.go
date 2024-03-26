package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Flushes stdout.
func FlushStdout() {
	writer := bufio.NewWriter(os.Stdout)
	writer.Flush()
}

// Prompts the user with a given string and returns their input.
func UserInput(s string) string {
	fmt.Print(s)
	FlushStdout()
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimSuffix(userInput, "\n")
}

func main() {
	name := UserInput("What is your name? ")
	fmt.Printf("Hello %s!\n", name)
}
