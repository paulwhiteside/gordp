package rdp

import(
	"fmt"
)

type Interpreter struct{

}

func (interpreter *Interpreter) Run(node Node) interface{}{
	//fmt.Println("interpreter.Run")

	var rv interface{}



	switch node.(type){
		case NumberNode:
			rv = interpreter.evalNumberNode(node.(NumberNode))
	    case AddNode:
			rv = interpreter.evalAddNode(node.(AddNode))
		case MultiplyNode:
			rv = interpreter.evalMultiplyNode(node.(MultiplyNode))
	}

	return rv
}

func (interpreter *Interpreter) evalNumberNode(node NumberNode) interface{}{
	fmt.Println("evalNumberNode")
	return node.value
}

func (interpreter *Interpreter) evalAddNode(node AddNode) interface{}{
	fmt.Println("evalAddNode")
	left := interpreter.Run(node.left)
	_ = left
	right := interpreter.Run(node.right)
	_ = right
	fmt.Println(left, right)

	var result interface{}

	switch left.(type){
	    case int:
			switch right.(type){
				case int:
					result = left.(int) + right.(int)
					fmt.Println(">>>", result)
			}
		case float64:
		case string:
	}

    return result
}

func (interpreter *Interpreter) evalMultiplyNode(node Node) interface{}{
	fmt.Println("evalMultiplyNode")
    return 123
}

