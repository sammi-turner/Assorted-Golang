// The Shunting Yard Algorithm is an algorithm created by Edsger Dijkstra that transforms an infix expression into a postfix (Reverse Polish Notation) expression.

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(str string) {
	s.items = append(s.items, str)
}

func (s *Stack) Pop() string {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() string {
	if len(s.items) > 0 {
		return s.items[len(s.items)-1]
	}
	return ""
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	default:
		return -1
	}
}

func addSpaces(expression string) string {
	withSpaces := ""

	for _, ch := range expression {
		switch ch {
		case '+', '-', '*', '/', '(', ')', '^':
			withSpaces += " " + string(ch) + " "
		default:
			withSpaces += string(ch)
		}
	}

	return strings.Join(strings.Fields(withSpaces), " ")
}

func ShuntingYard(expression string) string {
	outputQueue := ""
	operatorStack := &Stack{}
	expression = addSpaces(expression)
	tokens := strings.Fields(expression)

	for _, token := range tokens {
		if _, err := strconv.ParseFloat(token, 64); err == nil {
			outputQueue += token + " "
		} else {
			switch token {
			case "(":
				operatorStack.Push(token)
			case ")":
				for !operatorStack.IsEmpty() && operatorStack.Peek() != "(" {
					outputQueue += operatorStack.Pop() + " "
				}
				operatorStack.Pop()
			default:
				for !operatorStack.IsEmpty() && precedence(token) <= precedence(operatorStack.Peek()) {
					outputQueue += operatorStack.Pop() + " "
				}
				operatorStack.Push(token)
			}
		}
	}

	for !operatorStack.IsEmpty() {
		outputQueue += operatorStack.Pop() + " "
	}

	return strings.TrimSpace(outputQueue)
}

func EvaluatePostfix(expression string) float64 {
	stack := &StackFloat{}
	tokens := strings.Fields(expression)

	for _, token := range tokens {
		if val, err := strconv.ParseFloat(token, 64); err == nil {
			stack.Push(val)
		} else {
			if stack.IsEmpty() {
				fmt.Println("Invalid postfix expression")
				return 0
			}
			right := stack.Pop()

			if stack.IsEmpty() {
				fmt.Println("Invalid postfix expression")
				return 0
			}
			left := stack.Pop()

			switch token {
			case "+":
				stack.Push(left + right)
			case "-":
				stack.Push(left - right)
			case "*":
				stack.Push(left * right)
			case "/":
				stack.Push(left / right)
			case "^":
				stack.Push(math.Pow(left, right))
			}
		}
	}

	if stack.IsEmpty() {
		fmt.Println("Invalid postfix expression")
		return 0
	}
	return stack.Pop()
}

type StackFloat struct {
	items []float64
}

func (s *StackFloat) Push(f float64) {
	s.items = append(s.items, f)
}

func (s *StackFloat) Pop() float64 {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *StackFloat) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	postfix := ShuntingYard("(3+4.2)*2")
	fmt.Println(postfix)
	fmt.Println(EvaluatePostfix(postfix))
}
