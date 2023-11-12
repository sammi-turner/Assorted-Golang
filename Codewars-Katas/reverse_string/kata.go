package main

import "fmt"

func Solution(word string) string {
	var result string
	for _,v := range word {
		result = string(v) + result
	}
	return result
}

func main() {
	fmt.Println(Solution("hello world"))
}
