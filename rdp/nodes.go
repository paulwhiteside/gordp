package rdp

type Node interface{
    visit()
}

type NumberNode struct{
	value interface{}
}

func (numbernode NumberNode) visit(){
}

type AddNode struct{
	left Node
	right Node
}

func (addnode AddNode) visit(){
}

type SubtractNode struct{
	left Node
	right Node	
}

func (subtractnode SubtractNode) visit(){
}

type MultiplyNode struct{
	left Node
	right Node
}

func (multiplynode MultiplyNode) visit(){
}

type DivideNode struct{
	left Node
	right Node
}

func (dividenode DivideNode) visit(){
}

type PlusNode struct{
	node Node
}

func (plusnode PlusNode) visit(){
}

type MinusNode struct{
	node Node
}

func (minusnode MinusNode) visit() {
}

type ExponentNode struct{
	left Node
	right Node
}

type FunctionNode struct{

}

