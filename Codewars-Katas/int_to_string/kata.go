package main

import (
    "fmt"
    "strconv"
)

func NumberToString(n int) string {
    return strconv.Itoa(n)
}

func main() {
    r := NumberToString(29817)
    fmt.Println(r)
}
