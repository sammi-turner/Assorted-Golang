package main

import (
	"fmt"
	"unicode"
)

type MyString string

func (s MyString) IsUpperCase() bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) { return false }
	}
	return true
}

func main() {
	a := MyString.IsUpperCase("Is this all uppercase?")	
	b := MyString.IsUpperCase("IS THIS UPPERCASE?")
	fmt.Printf("%t %t\n", a, b)	
}
