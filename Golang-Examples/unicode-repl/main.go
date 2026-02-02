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

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("\nUnicode REPL - enter a character, decimal or 0xHex (empty string to exit)")
    for {
        fmt.Print("\n> ")
        if !scanner.Scan() {
            break
        }
        input := strings.TrimSpace(scanner.Text())
		if input == "" {
			fmt.Println("\n")
            return
		}
    	u, err := strconv.ParseUint(input, 0, 32)
		switch err {
		case nil:
            r := rune(u)
            switch {
            case !utf8.ValidRune(r):
                fmt.Printf("%d (0x%X) is not a valid Unicode code point.\n", u, u)
            case r == ' ':
                fmt.Printf("That is the space character\n")
            case unicode.IsGraphic(r):
                fmt.Printf("%c\n", r)
            default:
                fmt.Printf("Non-graphical character: %q\n", r)
            }
        default:
            fmt.Println("Please enter a number or hex value.")
        }
    }
}
