package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Prompts the user with a given string and returns their input.
func UserInput(s string) string {
	fmt.Print(s)
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimSuffix(userInput, "\n")
}

func main() {
	first := UserInput("What is first your name? ")
	last := UserInput("What is your last name? ")
	fmt.Printf("Hello %s %s!\n", first, last)
}
