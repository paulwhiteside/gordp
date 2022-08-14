package rdp

import(
	_ "fmt"
)

type Interpreter struct{

}

func (interpreter *Interpreter) Run(node Node) interface{}{
	return node.visit()
}