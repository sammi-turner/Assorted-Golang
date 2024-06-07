package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"monkey-compiler/compiler"
	"monkey-compiler/eval"
	"monkey-compiler/lexer"
	"monkey-compiler/object"
	"monkey-compiler/parser"
	"monkey-compiler/repl"
	"monkey-compiler/vm"
)

func main() {
	// Start Monkey REPL
	if len(os.Args) == 1 {
		fmt.Println("This is the Monkey programming language!")
		fmt.Println("Feel free to type in commands")
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	// Run a Monkey script
	if err := runScript(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runScript(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read %s: %v", filename, err)
	}

	p := parser.New(lexer.New(string(data)))
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return errors.New(strings.Join(p.Errors(), "\n"))
	}

	// Process macros
	macroEnv := object.NewEnvironment()
	eval.DefineMacros(program, macroEnv)
	expanded := eval.ExpandMacros(program, macroEnv)

	// Compile the AST to bytecode
	c := compiler.New()
	if err := c.Compile(expanded); err != nil {
		return fmt.Errorf("Woops! Compilation failed: %s", err)
	}

	// Run bytecode instructions
	machine := vm.New(c.Bytecode())
	if err := machine.Run(); err != nil {
		return fmt.Errorf("Woops! Executing bytecode failed: %s", err)
	}

	return nil
}
