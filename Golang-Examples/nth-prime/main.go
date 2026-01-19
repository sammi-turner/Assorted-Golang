package main

import (
	"os"
	"errors"
	"fmt"
	"math"
	"bufio"
	"strings"
	"strconv"
)

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func userInput() string {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	return strings.TrimSuffix(s, "\n")
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nthPrime(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Input must be a positive integer.")
	}
	count := 0
	num := 1
	for count < n {
		num++
		if isPrime(num) {
			count++
		}
	}
	return num, nil
}

func main() {
	for {
		s := userInput()
		switch {
		case !isInt(s):
			return
		default:
			n, _ := strconv.Atoi(s)
			prime, err := nthPrime(n)
			switch err {
			case nil:
				fmt.Printf("%d is number %d in the list of primes.\n\n", prime, n)
			default:
				fmt.Println(err)
				return
			}
		}
	}
}
