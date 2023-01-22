package main

import (
	"fmt"
	"gordp/functions"
	"gordp/rdp"
)

func Go() int {
	return 1
}

func main() {
	//s := `3.14 + 9 / (-3 ^ 2) + 1`
	s := `1 + 2 + f_1(1,2-1.53663763,foo(10,20)) + 500`
	//s := `1+3 * 2-8 + 7 * 3 * 4+4.44`
	//s := `-2 + 1`

	ast := rdp.BuildAst(s)
	interpreter := rdp.NewIntrepreter()

	//Register functions with the interpreter
	interpreter.RegisterFunction("f_1", functions.Sum)
	interpreter.RegisterFunction("foo", functions.Sum)

	var result interface{}

	result = interpreter.Eval(ast)

	fmt.Println(s, "=", result)

	// TODO Move this round tripping to a proper unit test
	/*
		fmt.Println("---------------------------------------")
		n := 35
		for idx := 1; idx < n; idx++ {
			a := spreadsheet.ToBase26(idx)
			b := spreadsheet.FromBase26(a)
			msg := fmt.Sprintf("%2d %2s %2d", idx, a, b)
			fmt.Println(">>", msg)
		}

		mybook := spreadsheet.NewBook()
		mybook.AddSheet("Sheet 1", 10, 10)
		mybook.AddSheet("Sheet 2", 10, 10)
		mybook.AddSheet("Sheet 3", 10, 10)

		sheet2 := mybook.GetSheet("Sheet 2")

		sheet2.AddCell("A6", 3.14, "")
		fmt.Println("**>", sheet2.GetCell("A6"))
		cell := sheet2.GetCell("A6")

		cell.SetFormula("=2 * 7")
		cell.Calculate()
		fmt.Println("**>", sheet2.GetCell("A6"))
	*/
}
