package main

import "fmt"

func MakeRow(multiplier, size int) []int {
	i := 1
	var result []int
	for i <= size {
		result = append(result, i*multiplier)
		i++
	}
	return result
}

func MultiplicationTable(size int) [][]int {
	i := 1
	var table [][]int
	for i <= size {
		table = append(table, MakeRow(i, size))
		i++
	}
	return table
}

func main() {
	fmt.Printf("%v\n", MultiplicationTable(3))
}
