package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	secretNumber := rng.Intn(100) + 1

	fmt.Print("\nWelcome to the Number Guessing Game!\n\nI'm thinking of a number between 1 and 100...\n")

	for {
		fmt.Print("\nMake a guess? ")
		var guess int
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			continue
		}

		if guess < secretNumber {
			fmt.Print("\nToo low...\n")
		} else if guess > secretNumber {
			fmt.Print("\nToo high...\n")
		} else {
			fmt.Print("\nYou've guessed the number correctly!\n\n")
			break
		}
	}
}
