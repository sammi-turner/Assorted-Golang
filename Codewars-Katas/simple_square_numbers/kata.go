package main

import (
	"fmt"
	"math"
)

func Solve(n int) int {
	for i := 1; i <= n/2; i++ {
		sqrt := math.Sqrt(float64(n + i*i))
		if int(sqrt)*int(sqrt) == n+i*i {
			return i * i
		}
	}
	return -1
}

func main() {
	fmt.Printf("%d %d %d\n", Solve(3), Solve(4), Solve(9))
}
