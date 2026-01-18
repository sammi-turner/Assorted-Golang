package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)
	attempts := 0
	fmt.Println("\nGuess the secret number between 1 and 100...")
	for {
		attempts++
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		guess, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("\nPlease enter a valid integer.")
			continue
		}
		switch {
		case guess < secret:
			fmt.Println("\nToo low...")
		case guess > secret:
			fmt.Println("\nToo high...")
		default:
			fmt.Printf("\nCorrect! You guessed the number in %d attempts.\n\n", attempts)
			return
		}
	}
}
