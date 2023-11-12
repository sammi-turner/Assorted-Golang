package main

import "fmt"

func GetSum(a, b int) int {
	if a > b { a, b = b, a }
	return (a + b) * (b - a + 1) / 2
}

func main() {
	a := GetSum(0, 1)
	b := GetSum(2, 5)
	c := GetSum(-3, 0)
	fmt.Printf("%d %d %d\n", a, b, c)
}
