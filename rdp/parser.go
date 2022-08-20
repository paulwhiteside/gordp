// Parses the grammer:
// E = Expression, T = Term, F = Factor, P = Power
//
// E -> T + E | T - E | E
// T -> F * T | F / T | F
// F -> P | ID | Integer | (E) | -F
// P -> F ^ E

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


func (parser *Parser) incr() {
	if parser.index < len(parser.tokens)-1{
		parser.index++
	}
	fmt.Printf("incr index=%d current_token=%s\n", parser.index, parser.current_token())
}

func (parser *Parser) Parse() Node{
	result := parser.expr()
	return result
}

func (parser *Parser) expr() Node{
	var result Node
	result = parser.term()

	if parser.current_token().tokentype == Plus{
		left := result
		parser.incr()
		right := parser.term()
	    result = AddNode{"Add", left, right}
	}

	fmt.Println("returning ", result)
	return result
}

func (parser *Parser) term() Node{
	var result Node
	result = parser.factor()

	if parser.current_token().tokentype == Multiply{
		left := result
		parser.incr()
		right := parser.factor()
		result = MultiplyNode{"Multiply", left, right}
		parser.incr()
	}else if parser.current_token().tokentype == Divide{
		left := result
		parser.incr()
		right := parser.factor()
		result = DivideNode{"Divide", left, right}
		parser.incr()
	}

	fmt.Println("result->", result)
	return result
}

func (parser *Parser) factor() Node{
	var result Node
	
	if parser.current_token().tokentype == LParen{
		parser.incr()
		result = parser.expr()
		if parser.current_token().tokentype != RParen{
			fmt.Println("Error.  Expected ) got ", parser.current_token())
			os.Exit(0)
		}
		parser.incr()
	}else if parser.current_token().tokentype == Number{
		result = NumberNode{"Number", parser.current_token().value}
		parser.incr()
	}else if parser.current_token().tokentype == Plus {
		parser.incr()
		result = PlusNode{"Plus", parser.factor()}
	}else if parser.current_token().tokentype == Minus {
		parser.incr()
		result = MinusNode{"Minus", parser.factor()}
		
	}
	return result
}


/*
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
		result = MultiplyNode{Multiply, result, parser.factor()}	
		parser.index++
	}else if parser.current_token().tokentype == Divide{
		result = DivideNode{result, parser.factor()}
		parser.index++
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

*/