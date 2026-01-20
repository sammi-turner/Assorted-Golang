package main

import (
	"fmt"
)

func main() {
    r, err := ReadFile("test.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("%s\n", r)
}