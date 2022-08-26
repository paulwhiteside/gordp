package main

import (
	"fmt"
	"gordp/rdp"
)

func main() {
	//s := `3.14 + 9 / (-3 ^ 2) + 1`
	//s := `1 + 2 + 3 + 4 + 5 + 6`
	s := `1+3 * 2-8 + 7 * 3 * 4+4`
	tokens := rdp.Lexer(s)
	fmt.Println(tokens)
	parser := rdp.NewParser(0, tokens)
	tree := parser.Parse()
	fmt.Println(tree)
	interpreter := rdp.Interpreter{}

	fmt.Println("---------------------------------------")
	var result interface{}
	for i := 0; i < 1000000; i++ {
		result = interpreter.Eval(tree)
	}
	fmt.Println(s, "=", result)

}
