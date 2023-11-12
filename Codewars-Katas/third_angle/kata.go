package main

import "fmt"

func OtherAngle(a int, b int) int {
    return 180 - a - b
}

func main() {
	fmt.Println(OtherAngle(30, 60))
}
