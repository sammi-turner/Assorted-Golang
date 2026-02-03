package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "unicode"
    "unicode/utf8"
)

func UserInput(s string) string {
	fmt.Print(s)
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimSuffix(userInput, "\n")
}

func main() {
    fmt.Println("\nUnicode REPL - enter a character, decimal or 0xHex (empty string to exit)")
    for {
        input := UserInput("\n> ")
		if input == "" {
            fmt.Println()
            return
		}
    	u, err := strconv.ParseUint(input, 0, 32)
		switch err {
		case nil:
            r := rune(u)
            switch {
            case !utf8.ValidRune(r):
                fmt.Println("Not a valid rune.")
            case r == ' ':
                fmt.Println("Space character.")
            case unicode.IsGraphic(r):
                fmt.Printf("%c\n", r)
            default:
                fmt.Println("Does not correspond to a rune.")
            }
        default:
            fmt.Println("Please enter a number or hex value.")
        }
    }
}
