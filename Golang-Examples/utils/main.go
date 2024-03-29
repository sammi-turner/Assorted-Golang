package main

import (
	"fmt"
	utils "utils/lib"
)

func main() {
	n := 0
	rng := utils.NewRNG()
	fmt.Print("Here are five dice rolls : ")
	for n < 5 {
		num := rng.Intn(6) + 1
		fmt.Printf("%d ", num)
		n += 1
	}
	fmt.Println()
}
