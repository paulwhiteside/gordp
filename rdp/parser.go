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
	var current_token Token
	if parser.index >= len(parser.tokens){
		current_token = Token{}
	}else{
		current_token = parser.tokens[parser.index]
	}
	return current_token
}


func (parser *Parser) incr() {
	parser.index++
	fmt.Printf("incr index=%d current_token=%s\n", parser.index, parser.current_token())
}

func (parser *Parser) Parse() Node{
	result := parser.expr()
	return result
}

func (parser *Parser) expr() Node{
	var result Node
	result = parser.term()
	current_token := parser.current_token()
	if current_token.tokentype == Plus{
		left := result
		fmt.Print("A ")
		parser.incr()
		right := parser.term()
		result = AddNode{"Add", left, right}
	}

	fmt.Println("returning ", result)
	parser.incr()
	return result
}

func (parser *Parser) term() Node{
	var result Node
	result = parser.factor()

	current_token := parser.current_token()
	if current_token.tokentype == Multiply{
		left := result
		fmt.Print("B ")
		parser.incr()
		right := parser.factor()
		result = MultiplyNode{"Multiply", left, right}
		fmt.Print("C ")
		parser.incr()
	}else if current_token.tokentype == Divide{
		left := result
		fmt.Print("D ")
		parser.incr()
		right := parser.factor()
		result = DivideNode{"Divide", left, right}
		fmt.Print("E ")
		parser.incr()
	}


	fmt.Println("result->", result)
	
	return result
}

func (parser *Parser) factor() Node{

	var result Node
	result = parser.factor_base()

	current_token := parser.current_token()
	if current_token.tokentype == Exponent{
		parser.incr()
		result = ExponentNode{"Exponent", result, parser.term()}
	}

	return result

}

func (parser *Parser) factor_base() Node{
	var result Node
	
	current_token := parser.current_token()
	if current_token.tokentype == LParen{
		fmt.Print("F ")
		parser.incr()
		result = parser.expr()
		if current_token.tokentype != RParen{
			fmt.Println("Error.  Expected ) got ", parser.current_token())
			os.Exit(0)
		}
		fmt.Print("G ")
		parser.incr()
	}else if current_token.tokentype == Number{
		result = NumberNode{"Number", parser.current_token().value}
		fmt.Print("H ")
		parser.incr()
	}else if current_token.tokentype == Plus {
		fmt.Print("I ")
		parser.incr()
		result = PlusNode{"Plus", parser.factor()}
	}else if current_token.tokentype == Minus {
		fmt.Print("J ")
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