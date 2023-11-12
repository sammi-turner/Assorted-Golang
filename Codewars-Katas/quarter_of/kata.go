package main

import "fmt"

func QuarterOf(month int) int {
	return ((month - 1)/ 3) + 1
}

func main() {
	a := QuarterOf(1)
	b := QuarterOf(5)
	c := QuarterOf(9)
	fmt.Printf("%d %d %d\n", a, b, c)
}
