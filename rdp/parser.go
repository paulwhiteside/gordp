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
	"strconv"
	"strings"
)

type Parser struct {
	index  int
	tokens []Token
}

func NewParser(index int, tokens []Token) Parser {
	parser := Parser{index, tokens}
	return parser
}

func (parser *Parser) current_token() Token {
	var current_token Token
	if parser.index < len(parser.tokens) {
		current_token = parser.tokens[parser.index]
	}
	return current_token
}

func (parser *Parser) peek_token() Token {
	var peek_token Token
	if parser.index < len(parser.tokens)-1 {
		peek_token = parser.tokens[parser.index+1]
	}
	return peek_token
}

func (parser *Parser) incr() {
	parser.index++
}

func (parser *Parser) Parse() Node {
	result := parser.expr()
	return result
}

func (parser *Parser) expr() Node {
	var result Node

	result = parser.term()

	for parser.index < len(parser.tokens) {

		current_token := parser.current_token()

		if current_token.tokentype == Plus {
			left := result
			parser.incr()
			right := parser.term()
			result = AddNode{"Add", left, right}
		} else if current_token.tokentype == Minus {
			left := result
			parser.incr()
			right := parser.term()
			result = SubtractNode{"Subtract", left, right}
		} else {
			break
		}

	}

	return result
}

func (parser *Parser) term() Node {
	var result Node
	result = parser.factor()

	for parser.index < len(parser.tokens) {
		current_token := parser.current_token()
		if current_token.tokentype == Multiply {
			left := result
			parser.incr()
			right := parser.factor()
			result = MultiplyNode{"Multiply", left, right}
		} else if current_token.tokentype == Divide {
			left := result
			parser.incr()
			right := parser.factor()
			result = DivideNode{"Divide", left, right}
		} else {
			break
		}
	}

	return result
}

func (parser *Parser) factor() Node {

	var result Node

	if parser.current_token().tokentype == Identifier && parser.peek_token().tokentype == LParen {
		fmt.Println("found a function", parser.current_token().value)
		result = parser.function()
	} else {
		result = parser.factor_base()
	}

	current_token := parser.current_token()
	if current_token.tokentype == Exponent {
		parser.incr()
		result = ExponentNode{"Exponent", result, parser.term()}
	}

	return result

}

func (parser *Parser) factor_base() Node {
	var result Node

	current_token := parser.current_token()
	if current_token.tokentype == LParen {
		parser.incr()
		result = parser.expr()
		current_token = parser.current_token()
		if current_token.tokentype != RParen {
			os.Exit(0)
		}
		parser.incr()
	} else if current_token.tokentype == Number {
		v := parser.current_token().value
		i := strings.Index(v, ".")
		if i != -1 {
			v_float, _ := strconv.ParseFloat(v, 64)
			result = NumberNode{"Number", v_float}
		} else {
			v_int, _ := strconv.Atoi(v)
			result = NumberNode{"Number", v_int}
		}
		parser.incr()
	} else if current_token.tokentype == Plus {
		parser.incr()
		result = PlusNode{"Plus", parser.factor()}
	} else if current_token.tokentype == Minus {
		parser.incr()
		result = MinusNode{"Minus", parser.factor()}
	}
	return result
}

func (parser *Parser) function() Node {
	var func_args []interface{}
	_ = func_args

	func_name := parser.current_token().value
	parser.incr()

	for parser.index < len(parser.tokens) {
		parser.incr()
		if parser.current_token().tokentype != Comma {
			result := parser.expr()
			func_args = append(func_args, result)
		}
		fmt.Println("**>", parser.current_token())
		if parser.current_token().tokentype == RParen {
			parser.incr() //move past the RParen
			break
		}
	}
	fmt.Println("creating FunctionNode", func_name)
	return FunctionNode{"function", func_name, func_args}
}
