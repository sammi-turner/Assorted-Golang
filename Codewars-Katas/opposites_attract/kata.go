package main

import "fmt"

func LoveFunc(flower1, flower2 int) bool {
  return (flower1 + flower2) % 2 == 1
}

func main() {
	fmt.Println(LoveFunc(12, 13))
}
