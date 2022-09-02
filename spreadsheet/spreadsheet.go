package spreadsheet

import (
	"fmt"
	"gordp/rdp"
)

const (
	default_rows = 100
	default_cols = 40
)

type Cell struct {
	interpreter *rdp.Interpreter
	cellref     string
	formula     string
	value       interface{}
	ast         rdp.Node
}

func (cell Cell) String() string {
	return fmt.Sprintf("{%T ref=%s, formula=%s, value=%v}", cell, cell.cellref, cell.formula, cell.value)
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
}

func (cell *Cell) Calculate() interface{} {
	cell.value = cell.interpreter.Eval(cell.ast)
	return cell.value
}

func (cell *Cell) GetValue() interface{} {
	return 101
}

type Sheet struct {
	book  *Book
	name  string
	cells [][]*Cell
}

func (sheet Sheet) String() string {
	return fmt.Sprintf("{%T name=%s cells=[%d][%d]}", sheet, sheet.name, 2, 2)
}

func (sheet *Sheet) AddCell(cellref string, value interface{}, formula string) {
	cell := &Cell{}
	cell.interpreter = &sheet.book.interpreter
	cell.cellref = cellref
	cell.value = value
	cell.formula = formula

	row, column := CellRefToCoords(cellref)
	fmt.Println("row", row, "     col", column)
	sheet.cells[row][column] = cell
}

func (sheet *Sheet) GetCell(cellref string) *Cell {
	row, column := CellRefToCoords(cellref)
	return sheet.cells[row][column]
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

func (book *Book) AddSheet(name string, rows, columns int) {
	sheet := Sheet{}
	sheet.book = book
	sheet.name = name

	sheet.cells = make([][]*Cell, rows)
	for i := 0; i < columns; i++ {
		sheet.cells[i] = make([]*Cell, columns)
	}

	for row := 0; row < len(sheet.cells); row++ {
		currentRow := sheet.cells[row]
		for column := 0; column < len(currentRow); column++ {
			cell := &Cell{}
			cell.cellref = CellRefFromCoords(column, row)
			sheet.cells[row][column] = cell
		}
	}

	book.sheets[name] = sheet
}
