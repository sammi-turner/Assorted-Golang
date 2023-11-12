package main

import "fmt"

func IsValidWalk(walk []rune) bool {
	x := 0
	y := 0
	i := 0

	l := len(walk)
	if l != 10 {
		return false
	}

	for i < l {
		switch walk[i] {
		case 'w':
			x--
		case 'e':
			x++
		case 's':
			y--
		case 'n':
			y++
		}
		i++
	}

	return x == 0 && y == 0
}

func main() {
	t0 := []rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}
	t1 := []rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}
	t2 := []rune{'w'}
	t3 := []rune{'n', 'n', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}
	fmt.Printf("%t %t %t %t\n", IsValidWalk(t0), IsValidWalk(t1), IsValidWalk(t2), IsValidWalk(t3))
}
