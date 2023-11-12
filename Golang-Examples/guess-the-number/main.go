package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    secretNumber := rand.Intn(100) + 1

    fmt.Println("Welcome to the Guess the Number Game!\n\nI'm thinking of a number between 1 and 100...")

    for {
        fmt.Print("\nMake a guess? ")
        var guess int
        _, err := fmt.Scan(&guess)
        
        if err != nil {
            fmt.Println("Invalid input! Please enter a number.")
            continue
        }

        if guess < secretNumber {
            fmt.Println("\nToo low...")
        } else if guess > secretNumber {
            fmt.Println("\nToo high...")
        } else {
            fmt.Println("\nYou've guessed the number correctly!\n")
            break
        }
    }
}
