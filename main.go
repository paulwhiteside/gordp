package main

import (
	"fmt"
	"strconv"
	"regexp"
	"gordp/rdp"
)


func main() {
	//s := `123.111 + (456 * 22) & "Fred is a ""clampet"""`
	s := `(3.14 + 9) * 3`

	tokens := rdp.Lexer(s)
	fmt.Println(tokens)
	parser := rdp.NewParser(0, tokens)
	tree := parser.Parse()
	fmt.Println(tree)

	fmt.Println("---------------------")
	s = `3.14 + 9 * 3`

	tokens = rdp.Lexer(s)
	fmt.Println(tokens)
	parser = rdp.NewParser(0, tokens)
	tree = parser.Parse()
	fmt.Println(tree)

	interpreter := rdp.Interpreter{}
	result := interpreter.Run(tree)
	fmt.Println(result)

	fmt.Println("-------------------------")

	x, _ := strconv.ParseFloat("3",64)
	_ = x
	var i interface{} = x //6643573564579999999
	switch i.(type){
	    case int:
		    fmt.Println("int", i)
	    case float64:
	        fmt.Println("float", i)
	}

	y  := "123"
    val, err := strconv.Atoi(y)
    if err != nil {
        fmt.Printf("Value %s is not a number\n", y)
    } else {
        fmt.Printf("Value %s is a number with value %d\n", y, val)
    }

	m, mx := regexp.Match(`[0-9]+\.[0-9]`, []byte("3.14"))
	fmt.Println(">>>", m, mx)

	re := regexp.MustCompile(`r[ua]n`)

	ba := []byte("ran run run abc def abcd")
    fmt.Println(re.FindAllIndex(ba, -1)) // [[0 3] [4 7] [8 11]]

}
