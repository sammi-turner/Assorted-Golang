package main

import (
	"fmt"
)

func main() {
    r, err := ReadFile("test.txt")
    switch err {
    case nil:
        fmt.Printf("%s\n", r)
    default:
        fmt.Println("Error:", err)
        return
    }
}