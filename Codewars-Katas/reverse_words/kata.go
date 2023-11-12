package main

import "fmt"

func Backwards(str string) string {
	var result string	
	for _, v := range str {
    	result = string(v) + result
	}

	return result
}

func ReverseWords(str string) string {
	var current_word string
	var result string

	rune_array := []rune(str)
	l := len(rune_array)
	i := 0
	
	for i < l {
		if rune_array[i] == ' ' {
			if current_word != "" {
				result += Backwards(current_word)
				current_word = ""
			}
			result += " "
		} else {
			current_word += string(rune_array[i])
		}
		i++	
	}

	if current_word != "" {
		result += Backwards(current_word)
	}
	
	return string(result)
}

func main() {
	fmt.Printf("%s\n", ReverseWords("The quick brown fox jumps over the lazy dog."))
	fmt.Printf("%s\n", ReverseWords("double  spaced  words"))
}
