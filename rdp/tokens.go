package rdp

type TokenType string

const (
	Number     TokenType = "Number"
	Plus       TokenType = "Plus"
	Minus      TokenType = "Minus"
	Multiply   TokenType = "Multiply"
	Divide     TokenType = "Divide"
	Identifier TokenType = "Identifier"
	String     TokenType = "String"
	LParen     TokenType = "LParen"
	RParen     TokenType = "RParen"
	Equals     TokenType = "Equals"
	Ampersand  TokenType = "Ampersand"
	Exponent   TokenType = "Exponent"
	Comma      TokenType = "Comma"
)

type Token struct {
	tokentype TokenType
	value     string
}
