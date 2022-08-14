package rdp

import (
	"fmt"
	"os"
)

type Parser struct {
	index int
	tokens[]Token
}

func NewParser(index int, tokens []Token) *Parser{
	parser := Parser{index, tokens}
	return &parser
}

func (parser *Parser) current_token () Token{
    return parser.tokens[parser.index]
}

func (parser *Parser) Parse() Node{
    fmt.Println("-- Parse --------------------")
	result := parser.expr()
	return result
}

func (parser *Parser) expr() Node{
	fmt.Println("-- expr ---------------------")
	fmt.Println("$$$$ expr current_token", parser.current_token())
	result := parser.term()
	fmt.Println("$$$$ expr current_token", parser.current_token())
	
	
	for parser.current_token().tokentype == Plus || parser.current_token().tokentype == Minus{
		if parser.current_token().tokentype == Plus{
			parser.index++
			result = AddNode{result, parser.term()}
		}
		if parser.current_token().tokentype == Minus{
			parser.index++
			result = SubtractNode{result, parser.term()}
		}
	}
	
	return result
}

func (parser *Parser) term() Node{
	fmt.Println("-- term ---------------------")
	fmt.Println("$$$$ term current_token", parser.current_token())
	result := parser.factor()
	parser.index++
	fmt.Println("$$$$ term current_token", parser.current_token())

	for parser.current_token().tokentype == Multiply || parser.current_token().tokentype == Divide {
		if parser.current_token().tokentype == Multiply{
			parser.index++
			result = MultiplyNode{result, parser.factor()}
			
		}
		if parser.current_token().tokentype == Divide{
			parser.index++
			result = DivideNode{result, parser.factor()}
			fmt.Println(result)
		}
	}
	return result
}

func (parser *Parser) factor() Node{
	fmt.Println("-- factor ---------------------")
	fmt.Println("$$$$ factor current_token", parser.current_token())
	result := parser.factor_base()
	parser.index++
	fmt.Println("$$$$ factor current_token", parser.current_token())
	return result
}

func (parser *Parser) factor_base() Node{
	fmt.Println("-- factor base ---------------------")
	fmt.Println("$$$$ factor_base current_token", parser.current_token())
	var result Node

	if parser.current_token().tokentype == LParen{
		parser.index++
		result = parser.term()
		if parser.current_token().tokentype != RParen{
			fmt.Println("Error")
			os.Exit(0)
		}
		parser.index++
	}else if parser.current_token().tokentype == Number{
		parser.index++
		result = NumberNode{parser.current_token().value}
	}else if parser.current_token().tokentype == Plus {
		parser.index++
		result = PlusNode{parser.factor()}
	}else if parser.current_token().tokentype == Minus {
		parser.index++
		result = MinusNode{parser.factor()}
	}else{
		fmt.Println("Error")
		os.Exit(0)
	}
	fmt.Println("$$$$ factor_base current_token", parser.current_token())
	return result
}