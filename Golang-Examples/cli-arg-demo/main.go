package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	switch len(args) {
	case 1:
		fmt.Println("No argument supplied.")
	case 2:
		fmt.Print("You entered: ")
		fmt.Println(args[1])
	default:
		fmt.Println("Only one argument is allowed.")
	}
}
