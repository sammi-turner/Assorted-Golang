package main

import (
	"fmt"
	"strings"
)

func DoPart(arr []string, comma int) string {
	l := len(arr)
	i := 0
	res := ""

	for i < l {
		res += arr[i]
		if i == comma {
			res += ", "
		} else {
			res += " "
		}
		i++
	}

	res = strings.TrimSuffix(res, " ")
	res = strings.TrimSuffix(res, ",")
	return "(" + res + ")"
}

func PartList(arr []string) string {
	l := len(arr) - 1
	i := 0
	res := ""

	for i < l {
		res += DoPart(arr, i)
		i++
	}

	return res
}

func main() {
	test := []string{"I", "wish", "I", "hadn't", "come"}	
	fmt.Printf("%s\n", PartList(test))
}
