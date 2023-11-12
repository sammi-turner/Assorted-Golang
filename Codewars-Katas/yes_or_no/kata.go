package main

import "fmt"

func BoolToWord(word bool) string {
	if word { return "Yes" }
	return "No"
}

func main() {
	value_t := true
	value_f := false
	fmt.Printf("%s %s\n", BoolToWord(value_t), BoolToWord(value_f))	
}
