package spreadsheet

import (
	"fmt"
	"gordp/rdp"
)

type Cell struct {
	interpreter *rdp.Interpreter
	cellref     string
	formula     string
	value       interface{}
	ast         rdp.Node
}

func NewCell(cellref string, value interface{}) *Cell {
	cell := Cell{}
	cell.cellref = cellref
	cell.value = value
	return &cell
}

func (cell Cell) String() string {
	return fmt.Sprintf("{%T %s, formula=%s, value=%v}", cell, cell.cellref, cell.formula, cell.value)
}

func (cell *Cell) SetFormula(formula string) {
	cell.value = nil //set any existing value to nil as formula will re-calculate it
	cell.formula = formula
	var tokens []rdp.Token
	if formula[0] == '=' {
		tokens = rdp.Lexer(formula[1:])
	} else {
		tokens = rdp.Lexer(formula)
	}
	parser := rdp.NewParser(tokens)
	cell.ast = parser.Parse()
	cell.value = cell.interpreter.Eval(cell.ast)
}

func (cell *Cell) GetValue() interface{} {
	return 101
}

type Sheet struct {
	book  *Book
	name  string
	cells [2][2]Cell
}

func (sheet Sheet) String() string {
	return fmt.Sprintf("{%T name=%s cells=[%d][%d]}", sheet, sheet.name, 2, 2)
}

func (sheet *Sheet) AddCell(cellref string, value interface{}, formula string) *Cell {
	cell := Cell{}
	cell.interpreter = &sheet.book.interpreter
	cell.cellref = cellref
	cell.value = value
	cell.formula = formula

	return &cell
}

func (sheet *Sheet) GetCell(cellref string) *Cell {
	x, y := ToCoords("A1")
	return &sheet.cells[x][y]
}

type Book struct {
	interpreter rdp.Interpreter
	sheets      map[string]Sheet
}

func NewBook() *Book {
	var sheets = make(map[string]Sheet, 0)
	book := Book{}
	book.sheets = sheets
	book.interpreter = rdp.NewIntrepreter()
	return &book
}

func (book Book) String() string {
	return fmt.Sprintf("{%T sheets=[%d]}", book, len(book.sheets))
}

func (book *Book) GetSheet(name string) Sheet {
	return book.sheets[name]
}

func (book *Book) AddSheet(name string) {
	sheet := Sheet{}
	sheet.book = book
	sheet.name = name
	book.sheets[name] = sheet
}

func Test() string {
	return "spreadsheet.test"
}
