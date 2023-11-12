package main

import (
	"fmt"
	"strings"
)

func duplicate(r rune, n int) string {
	var result []rune
	ascii := int(r)
	if ascii > 90 { ascii -= 32 }
	
	result = append(result, rune(ascii))
	ascii += 32
	
	i := 0
	a := rune(ascii)
	for i < n {
		result = append(result, a) 
		i++
	}
	return string(append(result, rune('-')))
}

func Accum(s string) string {
	result := ""
    max := len(s)
	
	i := 0
	for i < max {
		r := rune(s[i])
		result += duplicate(r, i)
		i++ 	
	}
	return strings.TrimSuffix(result, "-")
}
	
func main() {
	fmt.Printf("%s\n%s\n%s\n", Accum("abcd"), Accum("RqaEzty"), Accum("cwAt"))
}
