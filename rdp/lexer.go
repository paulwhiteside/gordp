package rdp

import(
	"fmt"
	"strings"
	"strconv"
)


func Lexer(expr string) []Token{

	i := 0
	var c byte
	//var nc byte

	var char_buffer []byte
	var tokens []Token
	var token_type TokenType
	len_expr := len(expr) 
	for i < len_expr{
		c = expr[i]
		if (c >= '0' && c <= '9') || c == '.' {
			token_type = Number
			char_buffer = append(char_buffer, c)
		}else if (c >= 'A'  && c <= 'Z') || (c >= 'a' && c <= 'z' ) || c == '_'{
			token_type = Identifier
			char_buffer = append(char_buffer, c)

			for i < len_expr-1{
				i++
				c = expr[i]
				isAZ := (c >= 'A' && c <= 'Z')
				isaz := (c >= 'a' && c <= 'z')
				is_underscore := (c == '_')
				is_number := (c >= '0' && c <= '9')
				if isAZ || isaz || is_underscore || is_number{
					char_buffer = append(char_buffer, c)
				}else{
					i--
					break
				}
			}
		}else if c == '"'{
			token_type = String
			char_buffer = append(char_buffer, c)
			for i < len_expr-1{
				i++
				c = expr[i]
				if c == '"'{
					char_buffer = append(char_buffer, c)
					break
				}
				char_buffer = append(char_buffer, c)
			}
		}else{
			if len(char_buffer) > 0{
				v := string(char_buffer[:])
				char_buffer = nil

				fmt.Println("appending:", v)

				i := strings.Index(v, ".")
				if i != -1{
					fmt.Println("appending as float")
					v_float, _ := strconv.ParseFloat(v,64)
				    tokens = append(tokens, Token{tokentype: token_type, value: v_float})
				}else{
					fmt.Println("appending as int")
					v_int, _ := strconv.Atoi(v)
					tokens = append(tokens, Token{tokentype: token_type, value: v_int})
				}
				
			}

			if c == '+'{
				tokens = append(tokens, Token{tokentype: Plus, value: string(c)})
			}else if c == '-'{
				tokens = append(tokens, Token{tokentype: Minus, value: string(c)})
			}else if c == '*'{
				tokens = append(tokens, Token{tokentype: Multiply, value: string(c)})
			}else if c == '/'{
				tokens = append(tokens, Token{tokentype: Divide, value: string(c)})
			}else if c == '&'{
				tokens = append(tokens, Token{tokentype: Ampersand, value: string(c)})
			}else if c == '('{
				tokens = append(tokens, Token{tokentype: LParen, value: string(c)})
			}else if c == ')'{
				tokens = append(tokens, Token{tokentype: RParen, value: string(c)})
			}
		}

		i++
	}
	//process remaining characters in buffer
	v := string(char_buffer[:])
	if token_type == String{
		v = string(char_buffer[1:len(char_buffer)-1])
		v = strings.Replace(v, `""`, `"`, -1)
	}
	char_buffer = nil
	tokens = append(tokens, Token{tokentype: token_type, value: v})

	return tokens
}