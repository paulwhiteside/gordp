package main

import (
	"fmt"
	"gordp/rdp"
)

func main() {
	s := `3.14 + 9 / (-3 ^ 2)`

	tokens := rdp.Lexer(s)
	fmt.Println(tokens)
	parser := rdp.NewParser(0, tokens)
	tree := parser.Parse()
	fmt.Println(tree)

	interpreter := rdp.Interpreter{}
	fmt.Println("||||||||||||||||||||||||||")
	result := interpreter.Eval(tree)
	fmt.Println(s, "=", result)

}
