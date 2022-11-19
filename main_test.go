package main

import (
	"gordp/functions"
	"gordp/rdp"
	"testing"
)

func TestCalulate(t *testing.T) {
	s := `=1 + 2 + 3 + 4 + 5 + 6  + sum(1,2-1.53663763,3.14) + 500`
	ast := rdp.BuildAst(s)
	interpreter := rdp.NewIntrepreter()

	//Register functions with the interpreter
	interpreter.RegisterFunction("sum", functions.Sum)
	result := interpreter.Eval(ast)

	if result != 525.60336237 {
		t.Error("Error")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{`2+1`, 3},
		{`-2+1`, -1},
	}

	for _, test := range tests {
		ast := rdp.BuildAst(test.input)
		interpreter := rdp.NewIntrepreter()

		result := interpreter.Eval(ast)

		if result != test.expected {
			t.Error("input:{} expected:{} result:{}", test.input, test.expected, result)
		}

	}
}
