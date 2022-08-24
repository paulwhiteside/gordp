package main

import (
	"fmt"
	"gordp/rdp"
)

func main() {
	s := `3.14 + 9 / (-3 ^ 2)`

	tokens := rdp.Lexer(s)
	parser := rdp.NewParser(0, tokens)
	tree := parser.Parse()

	interpreter := rdp.Interpreter{}
	result := interpreter.Eval(tree)
	fmt.Println(s, "=", result)

}
