package main

import (
	"errors"
	"fmt"
	"math"
)

// Helper function to check if a number is prime
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

// Function to compute the nth prime, where 2 is the 1st.
func nthPrime(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("input must be a positive integer")
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
	prime, err := nthPrime(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(prime)
	}

	prime, err = nthPrime(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(prime)
	}
}
