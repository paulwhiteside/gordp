package rdp

func BuildAst(expr string) Node {
	var tokens []Token
	if expr[0] == '=' {
		tokens = Lexer(expr[1:])
	} else {
		tokens = Lexer(expr)
	}
	parser := NewParser(tokens)
	node := parser.Parse()
	return node
}
