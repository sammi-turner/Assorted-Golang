package main

import "fmt"

func countSheep(num int) string {
	i := 1
	s := ""
	
	for i <= num {
    	s += fmt.Sprintf("%d sheep...", i)
    	i++
	}
	
	return s
}

func main() {
	fmt.Println(countSheep(7))
}
