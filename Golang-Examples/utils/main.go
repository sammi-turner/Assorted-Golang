package main

import (
	"fmt"
	utils "utils/lib"
)

func main() {
	test := "Hello, 世界✨"
    size := utils.RuneLength(test)
	for i := 0; i < size; i++ {
		r, _ := utils.NthRune(test, i)
		fmt.Printf("%c\n", r)
	}
}
