package rdp

import (
	"fmt"
	"math"
	"os"
)

type Interpreter struct {
}

func (interpreter *Interpreter) Eval(node Node) interface{} {

	var rv interface{}

	switch node.(type) {
	case NumberNode:
		rv = interpreter.evalNumberNode(node.(NumberNode))
	case AddNode:
		rv = interpreter.evalAddNode(node.(AddNode))
	case SubtractNode:
		rv = interpreter.evalSubtractNode(node.(SubtractNode))
	case MultiplyNode:
		rv = interpreter.evalMultiplyNode(node.(MultiplyNode))
	case DivideNode:
		rv = interpreter.evalDivideNode(node.(DivideNode))
	case MinusNode:
		rv = interpreter.evalMinusNode(node.(MinusNode))
	case PlusNode:
		rv = interpreter.evalPlusNode(node.(PlusNode))
	case ExponentNode:
		rv = interpreter.evalExponentNode(node.(ExponentNode))
	default:
		fmt.Printf("SOME OTHER NODE %T\n", node)
		os.Exit(0)

	}

	return rv
}

func (interpreter *Interpreter) evalNumberNode(node NumberNode) interface{} {
	return node.value
}

func (interpreter *Interpreter) evalAddNode(node AddNode) interface{} {

	left := interpreter.Eval(node.left)
	right := interpreter.Eval(node.right)

	var result interface{}

	switch left.(type) {
	case int:
		switch right.(type) {
		case int:
			result = left.(int) + right.(int)
		case float64:
			result = left.(float64) + right.(float64)
		}
	case float64:
		switch right.(type) {
		case int:
			result = left.(float64) + float64(right.(int))
		case float64:
			result = left.(float64) + right.(float64)
		}
	}

	return result
}

func (interpreter *Interpreter) evalSubtractNode(node SubtractNode) interface{} {

	left := interpreter.Eval(node.left)
	right := interpreter.Eval(node.right)

	var result interface{}

	switch left.(type) {
	case int:
		switch right.(type) {
		case int:
			result = left.(int) - right.(int)
		case float64:
			result = left.(float64) - right.(float64)
		}
	case float64:
		switch right.(type) {
		case int:
			result = left.(float64) - float64(right.(int))
		case float64:
			result = left.(float64) - right.(float64)
		}
	}

	return result
}

func (interpreter *Interpreter) evalMultiplyNode(node MultiplyNode) interface{} {

	left := interpreter.Eval(node.left)
	right := interpreter.Eval(node.right)

	var result interface{}

	switch left.(type) {
	case int:
		switch right.(type) {
		case int:
			result = left.(int) * right.(int)
		case float64:
			result = float64(left.(int)) * right.(float64)
		}
	case float64:
		switch right.(type) {
		case int:
			result = left.(float64) * float64(right.(int))
		case float64:
			result = left.(float64) * right.(float64)
		}

	}

	return result
}

func (interpreter *Interpreter) evalDivideNode(node DivideNode) interface{} {

	left := interpreter.Eval(node.left)
	right := interpreter.Eval(node.right)

	var result interface{}

	switch left.(type) {
	case int:
		switch right.(type) {
		case int:
			result = left.(int) / right.(int)
		case float64:
			result = float64(left.(int)) / right.(float64)
		}
	case float64:
		switch right.(type) {
		case int:
			result = left.(float64) / float64(right.(int))
		case float64:
			result = left.(float64) / right.(float64)
		}

	}

	return result
}

func (interpreter *Interpreter) evalExponentNode(node ExponentNode) interface{} {

	left := interpreter.Eval(node.left)
	right := interpreter.Eval(node.right)

	var result interface{}

	switch left.(type) {
	case int:
		switch right.(type) {
		case int:
			result = math.Pow(float64(left.(int)), float64(right.(int)))
		case float64:
			result = math.Pow(left.(float64), right.(float64))
		}
	case float64:
		switch right.(type) {
		case int:
			result = math.Pow(left.(float64), float64(right.(int)))
		case float64:
			result = math.Pow(left.(float64), right.(float64))
		}

	}

	return result
}

func (interpreter *Interpreter) evalMinusNode(node MinusNode) interface{} {
	numbernode := interpreter.Eval(node.node)
	var result interface{}
	switch numbernode.(type) {
	case int:
		result = int(0) - numbernode.(int)
	case float64:
		result = float64(0.0) - numbernode.(float64)
	}
	return result
}

func (interpreter *Interpreter) evalPlusNode(node PlusNode) interface{} {
	numbernode := interpreter.Eval(node.node)
	var result interface{}
	switch numbernode.(type) {
	case int:
		result = int(0) + numbernode.(int)
	case float64:
		result = float64(0.0) + numbernode.(float64)
	}
	return result
}
