package rdp

import (
	"fmt"
	"os"
)

type Parser struct {
	index int
	tokens[]Token
}

func NewParser(index int, tokens []Token) Parser{
	parser := Parser{index, tokens}
	return parser
}

func (parser *Parser) current_token () Token{
	if parser.index >= len(parser.tokens){
		parser.index = len(parser.tokens)-1
	}
	return parser.tokens[parser.index]
}

/*
func (parser *Parser) advance() {
	if parser.index < len(parser.tokens)-1{
		parser.index++
	}
}
*/

func (parser *Parser) Parse() Node{
	result := parser.expr()
	return result
}

func (parser *Parser) expr() Node{
	result := parser.term()
	if parser.current_token().tokentype == Plus{
		result = AddNode{"Add",result, parser.term()}
	}else if parser.current_token().tokentype == Minus{
		result = SubtractNode{result, parser.term()}
	}
	return result
}

func (parser *Parser) term() Node{
	result := parser.factor()
	
	if parser.current_token().tokentype == Multiply{
		parser.index++
		result = MultiplyNode{Multiply, result, parser.factor()}	
	}else if parser.current_token().tokentype == Divide{
		parser.index++
		result = DivideNode{result, parser.factor()}
	}
	
	return result
}

func (parser *Parser) factor() Node{
	result := parser.factor_base()
	return result
}

func (parser *Parser) factor_base() Node{
	var result Node

	if parser.current_token().tokentype == LParen{
		parser.index++
		result = parser.expr()
		if parser.current_token().tokentype != RParen{
			fmt.Println("Error.  Expected ) got ", parser.current_token())
			os.Exit(0)
		}
		parser.index++
	}else if parser.current_token().tokentype == Number{
		result = NumberNode{parser.current_token().tokentype, parser.current_token().value}
		parser.index++
	}else if parser.current_token().tokentype == Plus {
		parser.index++
		result = PlusNode{parser.factor()}
	}else if parser.current_token().tokentype == Minus {
		parser.index++
		result = MinusNode{parser.factor()}
		
	}else{
		fmt.Println("Error", parser.current_token())
		os.Exit(0)
	}
	return result
}