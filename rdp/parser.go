// Parses the grammer:
// E = Expression, T = Term, F = Factor, P = Power
//
// E -> T + E | T - E | E
// T -> F * T | F / T | F
// F -> P | ID | Integer | (E) | -F
// P -> F ^ E

package rdp

import (
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
	if parser.index >= len(parser.tokens) {
		current_token = parser.tokens[len(parser.tokens)-1]
	} else {
		current_token = parser.tokens[parser.index]
	}
	return current_token
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
	current_token := parser.current_token()
	if current_token.tokentype == Plus {
		left := result
		parser.incr()
		right := parser.term()
		result = AddNode{"Add", left, right}
	}

	parser.incr()
	return result
}

func (parser *Parser) term() Node {
	var result Node
	result = parser.factor()

	current_token := parser.current_token()
	if current_token.tokentype == Multiply {
		left := result
		parser.incr()
		right := parser.factor()
		result = MultiplyNode{"Multiply", left, right}
		parser.incr()
	} else if current_token.tokentype == Divide {
		left := result
		parser.incr()
		right := parser.factor()
		result = DivideNode{"Divide", left, right}
		parser.incr()
	}

	return result
}

func (parser *Parser) factor() Node {

	var result Node
	result = parser.factor_base()

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
