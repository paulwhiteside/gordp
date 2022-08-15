package rdp

import (
	_ "fmt"
)

type Node interface{
}

type NumberNode struct{
	tokentype TokenType 
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
	tokentype TokenType
	left Node
	right Node
}

type DivideNode struct{
	left Node
	right Node
}

type PlusNode struct{
	node Node
}

type MinusNode struct{
	node Node
}

type ExponentNode struct{
	left Node
	right Node
}

type FunctionNode struct{

}

