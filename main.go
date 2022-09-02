package main

import (
	"fmt"
	"gordp/functions"
	"gordp/rdp"
	"gordp/spreadsheet"
)

func Go() int {
	return 1
}

func main() {
	//s := `3.14 + 9 / (-3 ^ 2) + 1`
	s := `1 + 2 + 3 + 4 + 5 + 6  + myfunc(1,2-1.53663763,3.14) + 500`
	//s := `1+3 * 2-8 + 7 * 3 * 4+4.44`
	tokens := rdp.Lexer(s)
	fmt.Println(tokens)
	parser := rdp.NewParser(tokens)
	ast := parser.Parse()
	fmt.Println(ast)
	interpreter := rdp.NewIntrepreter()

	//Register functions with the interpreter
	interpreter.RegisterFunction("myfunc", functions.Foo)

	fmt.Println("---------------------------------------")
	var result interface{}
	for i := 0; i < 1000000; i++ {
		result = interpreter.Eval(ast)
	}
	fmt.Println(s, "=", result)

	fmt.Println("---------------------------------------")
	for idx := 1; idx < 60; idx++ {
		a := spreadsheet.ToBase26(idx)
		b := spreadsheet.FromBase26(a)
		fmt.Println(">>", idx, a, b)
	}

	x, y := spreadsheet.ToCoords("A1")
	fmt.Println("-->", x, y)

	mybook := spreadsheet.NewBook()
	mybook.AddSheet("Sheet 1")
	mybook.AddSheet("Sheet 2")
	mybook.AddSheet("Sheet 3")

	sheet2 := mybook.GetSheet("Sheet 2")

	sheet2.AddCell("A1", 3.14, "")
	fmt.Println("**>", sheet2.GetCell("A1"))
	cell := sheet2.GetCell("A1")
	cell.SetFormula("=2 * 7")
	cell.Calculate()
	fmt.Println("**>", sheet2.GetCell("A1"))

}
