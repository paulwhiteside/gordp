package main

import (
	"fmt"
	"gordp/rdp"
)

func main() {
	//s := `3.14 + 9 / (-3 ^ 2) + 1`
	s := `1 + 2 + 3 + 4 + 5 + 6`
	//s := `1 * 2 * 3 * 4.7264726`
	tokens := rdp.Lexer(s)
	fmt.Println(tokens)
	parser := rdp.NewParser(0, tokens)
	tree := parser.Parse()
	fmt.Println(tree)
	interpreter := rdp.Interpreter{}

	fmt.Println("---------------------------------------")
	var result interface{}
	for i := 0; i < 10000000; i++ {
		result = interpreter.Eval(tree)
	}
	fmt.Println(s, "=", result)

}
