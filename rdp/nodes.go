package rdp

import (
	"fmt"
)

type Node interface{
    visit() interface{}
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

func (addnode AddNode) visit() interface{}{
	fmt.Println("Visiting AddNode")
	left := addnode.left.visit()
	right := addnode.right.visit()
	fmt.Println(left, "|", right)
	return 10001
}

type SubtractNode struct{
	left Node
	right Node	
}

func (subtractnode SubtractNode) visit() interface{}{
	return 1901
}

type MultiplyNode struct{
	tokentype TokenType
	left Node
	right Node
}

func (multiplynode MultiplyNode) visit() interface{}{
	return 4
}

type DivideNode struct{
	left Node
	right Node
}

func (dividenode DivideNode) visit() interface{}{
	return 3
}

type PlusNode struct{
	node Node
}

func (plusnode PlusNode) visit() interface{}{
	return 1
}

type MinusNode struct{
	node Node
}

func (minusnode MinusNode) visit() interface{}{
	return 2
}

type ExponentNode struct{
	left Node
	right Node
}

type FunctionNode struct{

}

