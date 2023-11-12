package main

import (
	"fmt"
	"strconv"
)

func StringToNumber(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil { panic(err) }
	return n
}

func main() {
	example := StringToNumber("1234")
	fmt.Printf("%d\n", example)
}
