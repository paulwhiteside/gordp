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
	if parser.index >= len(parser.tokens) {
		current_token = Token{}
	} else {
		current_token = parser.tokens[parser.index]
	}
	return current_token
}

func (parser *Parser) incr() {
	parser.index++
	fmt.Printf("incr index=%d current_token=%s\n", parser.index, parser.current_token())
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
		fmt.Print("A ")
		parser.incr()
		right := parser.term()
		result = AddNode{"Add", left, right}
	}

	fmt.Println("returning ", result)
	parser.incr()
	return result
}

func (parser *Parser) term() Node {
	var result Node
	result = parser.factor()

	current_token := parser.current_token()
	if current_token.tokentype == Multiply {
		left := result
		fmt.Print("B ")
		parser.incr()
		right := parser.factor()
		result = MultiplyNode{"Multiply", left, right}
		fmt.Print("C ")
		parser.incr()
	} else if current_token.tokentype == Divide {
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
		fmt.Print("F ")
		parser.incr()
		result = parser.expr()
		current_token = parser.current_token()
		if current_token.tokentype != RParen {
			fmt.Printf("Error.  Expected ) got %s      %T", current_token, current_token)
			os.Exit(0)
		}
		fmt.Print("G ")
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
		fmt.Print("H ")
		parser.incr()
	} else if current_token.tokentype == Plus {
		fmt.Print("I ")
		parser.incr()
		result = PlusNode{"Plus", parser.factor()}
	} else if current_token.tokentype == Minus {
		fmt.Print("J ")
		parser.incr()
		result = MinusNode{"Minus", parser.factor()}

	}
	return result
}
