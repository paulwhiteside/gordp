package main

import (
	"fmt"
	"gordp/functions"
	_ "gordp/functions"
	"gordp/rdp"
)

func main() {
	//s := `3.14 + 9 / (-3 ^ 2) + 1`
	s := `1 + 2 + 3 + 4 + 5 + 6  + myfunc(1,2-1.53663763,3.14) + 500`
	//s := `1+3 * 2-8 + 7 * 3 * 4+4.44`
	tokens := rdp.Lexer(s)
	fmt.Println(tokens)
	parser := rdp.NewParser(0, tokens)
	ast := parser.Parse()
	fmt.Println(ast)
	interpreter := rdp.NewIntrepreter()
	interpreter.RegisterFunction("myfunc", functions.Foo)

	fmt.Println("---------------------------------------")
	var result interface{}
	for i := 0; i < 1; i++ {
		result = interpreter.Eval(ast)
	}
	fmt.Println(s, "=", result)
}
