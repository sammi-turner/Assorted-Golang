package main

import (
	"fmt"
	utils "utils/lib"
)

func main() {
	name := utils.UserInput("What is your name? ")
	fmt.Printf("Hello %s!\n", name)
}
