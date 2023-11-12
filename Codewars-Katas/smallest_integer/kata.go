package main

import ( 
	"fmt"
	"sort"
)

func SmallestIntegerFinder(numbers []int) int {
  	sort.Ints(numbers)
  	return numbers[0]
}

func main() {
	arr := []int{34, -345, -1, 100}
	fmt.Println(SmallestIntegerFinder(arr))
}
