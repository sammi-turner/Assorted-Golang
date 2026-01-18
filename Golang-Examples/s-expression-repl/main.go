package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Token struct {
	Value string
	Type  string
}

func UserInput(s string) string {
	fmt.Print(s)
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimSuffix(userInput, "\n")
}

func TrimWhiteSpace(arg string) string {
	return strings.Join(strings.Fields(arg), " ")
}

func IsInt(arg string) bool {
	_, err := strconv.Atoi(arg)
	return err == nil
}

func IsFloat(arg string) bool {
	_, err := strconv.ParseFloat(arg, 64)
	return err == nil
}

func IsMantissaZero(arg string) bool {
	f, err := strconv.ParseFloat(arg, 64)
	switch err {
	case nil:
		return f == float64(int(f))
	default:
		return false
	}
}

func TruncateExcessZeros(arg string) string {
	result := strings.TrimRight(arg, "0")
	if strings.HasSuffix(result, ".") {
		result = result[:len(result)-1]
	}
	return result
}

func ToInt(arg string) int {
	n, _ := strconv.Atoi(arg)
	return n
}

func ToFloat(arg string) float64 {
	f, _ := strconv.ParseFloat(arg, 64)
	return f
}

func StringPosition(arg string, v []string) int {
	for i, s := range v {
		if arg == s {
			return i
		}
	}
	return -1
}

func SplitStringBySpace(arg string) []string {
	return strings.Fields(arg)
}

func RemoveEmptyStrings(arg []string) []string {
	var result []string
	for _, s := range arg {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}

func PadParens(arg string) string {
	var result string
	for _, c := range arg {
		switch {
		case c == '(':
			result += " ( "
		case c == ')':
			result += " ) "
		default:
			result += string(c)
		}
	}
	return result
}

func IsFunction(arg string) bool {
	v := []string{"add", "sub", "mul", "div", "mod", "pow"}
	return StringPosition(arg, v) > -1
}

func Tokenise(arg string) []Token {
	var t Token
	var result []Token
	p := PadParens(arg)
	v0 := SplitStringBySpace(p)
	v1 := RemoveEmptyStrings(v0)
	for _, s := range v1 {
		t.Value = s
		switch {
		case s == "(":
			t.Type = "OPEN"
		case s == ")":
			t.Type = "CLOSE"
		case IsFunction(s):
			t.Type = "FUNC"
		case IsInt(s) || IsFloat(s):
			t.Type = "NUMBER"
		default:
			t.Type = "ILLEGAL"
		}
		result = append(result, t)
	}
	return result
}

func Parse(arg []Token) string {
	openParen := 0
	size := len(arg)
	expected := []string{"OPEN", "NUMBER"}
	for i := 0; i < size; i++ {
		switch {
		case StringPosition(arg[i].Type, expected) == -1:
			return fmt.Sprintf("Unexpected token '%s' at position %d.", arg[i].Value, i)
		case arg[i].Type == "OPEN":
			openParen++
			expected = []string{"FUNC", "OPEN", "NUMBER"}
		case arg[i].Type == "CLOSE":
			openParen--
			switch openParen {
			case 0:
				expected = []string{"OPEN", "NUMBER"}
			default:
				expected = []string{"CLOSE", "OPEN", "NUMBER"}
			}
		case arg[i].Type == "NUMBER":
			expected = []string{"CLOSE", "OPEN", "NUMBER"}
		default:
			expected = []string{"OPEN", "NUMBER"}
		}
	}
	switch openParen {
	case 0:
		return "valid"
	default:
		return "Mismatched parentheses."
	}
}

func Evaluate(arg []Token) string {
	var funcs, operands []string
	for _, t := range arg {
		switch t.Type {
		case "FUNC":
			funcs = append(funcs, t.Value)
		case "NUMBER":
			operands = append(operands, t.Value)
		case "CLOSE":
			var n0, n1, n2 float64
			switch {
			case len(funcs) == 0:
				return "Parse error: too many operands."
			case len(operands) < 2:
				return "Parse error: too few operands."
			}
			f := funcs[len(funcs)-1]
			funcs = funcs[:len(funcs)-1]
			n1 = ToFloat(operands[len(operands)-1])
			operands = operands[:len(operands)-1]
			n0 = ToFloat(operands[len(operands)-1])
			operands = operands[:len(operands)-1]
			switch f {
			case "add":
				n2 = n0 + n1
			case "sub":
				n2 = n0 - n1
			case "mul":
				n2 = n0 * n1
			case "div":
				n2 = n0 / n1
			case "mod":
				n2 = float64(int(n0) % int(n1))
			case "pow":
				n2 = math.Pow(n0, n1)
			}
			operands = append(operands, fmt.Sprintf("%v", n2))
		}
	}
	switch {
	case len(funcs) != 0:
		return "Parse error: too few operands."
	case len(operands) > 1:
		return "Parse error: too many operands."
	case IsMantissaZero(operands[0]):
		r := ToInt(operands[0])
		return fmt.Sprintf("%d", r)
	default:
		return TruncateExcessZeros(operands[0])
	}
}

func main() {
	var s, p, r string
	var stream []Token
	fmt.Println("\nWelcome to the S-Expression REPL!")
	for {
		s = UserInput("\n> ")
		if s == "" {
			fmt.Println()
			break
		}
		stream = Tokenise(s)
		p = Parse(stream)
		switch p {
		case "valid":
			r = Evaluate(stream)
			fmt.Println(r)
		default:
			fmt.Println(p)
		}
	}
}
