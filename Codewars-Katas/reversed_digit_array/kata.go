package main

import "fmt"

func Digitize(n int) []int {
	var r []int
	for {
		r = append(r, n % 10)
		n /= 10
		if n == 0 {
			return r
		}
	}
}

func main() {
	res := Digitize(34567)
	fmt.Println(res)
}
