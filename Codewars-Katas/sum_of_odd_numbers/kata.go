package main

import "fmt"

func RowSumOddNumbers(n int) int {
    i := (n * n) - n + 1
	max := (n * n) + n - 1
	res := 0

	for i <= max {
		res += i
		i += 2
	}

	return res 
}

func main() {
	fmt.Printf("1 -> %d, 2 -> %d, 13 -> %d\n", RowSumOddNumbers(1), RowSumOddNumbers(2), RowSumOddNumbers(13))
}
