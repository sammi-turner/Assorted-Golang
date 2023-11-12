package main

import "fmt"

func GetGrade(a, b, c int) rune {
	n := (a + b + c) / 3
	if n >= 90 { return 'A' }
	if n >= 80 { return 'B' }
	if n >= 70 { return 'C' }
	if n >= 60 { return 'D' }
	return 'F'
}

func main() {
	average := GetGrade(65, 70, 59)
	fmt.Printf("%c\n", average)
}
