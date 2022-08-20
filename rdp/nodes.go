package rdp

import (
	_ "fmt"
)

type Node interface{
}



type NumberNode struct{
	nodetype string
	value interface{}
}

func (numbernode NumberNode) visit() interface{}{
	return numbernode.value
}

type AddNode struct{
	nodetype string
	left Node
	right Node
}

type SubtractNode struct{
	left Node
	right Node	
}

type MultiplyNode struct{
	nodetype string
	left Node
	right Node
}

type DivideNode struct{
	nodetype string
	left Node
	right Node
}

type PlusNode struct{
	nodetype string
	node Node
}

type MinusNode struct{
	nodetype string
	node Node
}

type ExponentNode struct{
	left Node
	right Node
}

type FunctionNode struct{

}

