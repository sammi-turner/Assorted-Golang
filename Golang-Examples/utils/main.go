package main

import (
	"fmt"
)

func main() {
    s := ReadFile("test.txt")
    fmt.Println(AppendString("Hello! ", s))
}
