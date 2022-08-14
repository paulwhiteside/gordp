package main

import (
	"fmt"
	"gordp/rdp"
)


func main() {
	//s := `123.111 + (456 * 22) & "Fred is a ""clampet"""`
	s := `(3.14 + 9) * 3`

	tokens := rdp.Lexer(s)
	fmt.Println(tokens)
	parser := rdp.NewParser(0, tokens)
	tree := parser.Parse()
	fmt.Println(tree)

	fmt.Println("---------------------")
	s = `3.14 + 9 * 3`

	tokens = rdp.Lexer(s)
	fmt.Println(tokens)
	parser = rdp.NewParser(0, tokens)
	tree = parser.Parse()
	fmt.Println(tree)

	interpreter := rdp.Interpreter{}
	result := interpreter.Run(tree)
	print(result)
}
