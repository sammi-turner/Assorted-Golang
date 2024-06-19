package main

import (
	"fmt"
	utils "utils/lib"
)

func main() {
	test := "Hello, 世界"
	n := 8

    r, err := utils.NthRune(test, n)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
		fmt.Printf("The rune at index %d is %c\n", n, r)
	}
}