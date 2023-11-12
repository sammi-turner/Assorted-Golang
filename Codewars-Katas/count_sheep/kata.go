package main

import "fmt"

func CountSheeps(numbers []bool) int {
	l := len(numbers)
	sheep := 0
	
	for i := 0; i < l; i++ {
		if numbers[i] == true { sheep++ }
	}
	
	return sheep
}

func main() {
	arr := []bool {true, true, true, false, true, true, true, true, true, false, true, false, true, false, false, true,
	true, true, true, true, false, false, true, true}
	fmt.Printf("%d\n", CountSheeps(arr))
}
