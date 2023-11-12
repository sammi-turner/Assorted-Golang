package main

import "fmt"

func MakeNegative(x int) int {
  	if x > 0 { return -x }
	return x
}

func main() {
	a := MakeNegative(3)
	b := MakeNegative(-3)
	c := MakeNegative(0)
	fmt.Printf("%d %d %d\n", a, b, c)
}
