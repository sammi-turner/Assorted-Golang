package main

import "fmt"

func Divisors(n int) int {
	i := 2
	count := 1
	
	for i <= n { 
		if n % i == 0 { count++ }
		i++ 
	}

	return count
}

func main() {
	fmt.Printf("%d %d %d %d\n", Divisors(4), Divisors(5), Divisors(12), Divisors(30))
}
