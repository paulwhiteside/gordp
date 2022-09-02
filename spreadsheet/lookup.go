package spreadsheet

import (
	"fmt"
	"strconv"
)

func ToBase26(colnum int) string {
	var dividend int
	var modulo int
	var b26 string

	dividend = colnum

	for dividend > 0 {
		modulo = (dividend - 1) % 26
		b26 = fmt.Sprintf("%c", (65+modulo)) + b26
		dividend = int((dividend - modulo) / 26)
	}

	return b26
}

func FromBase26(snum string) int {
	s := 0
	for i := 0; i < len(snum); i++ {
		if i == 0 {
			s = int(snum[i]) - 64
		} else {
			s *= 26
			s += int(snum[i]) - 64
		}
	}

	return s
}

func CellRefToCoords(cellref string) (int, int) {

	colref := ""
	rowref := ""

	i := 0
	for ; i < len(cellref); i++ {
		c := cellref[i]
		if c >= 'A' && c <= 'Z' {
			colref += fmt.Sprintf("%c", c)
		} else {
			break
		}
	}

	for ; i < len(cellref); i++ {
		c := cellref[i]
		if c >= '0' && c <= '9' {
			rowref += fmt.Sprintf("%c", c)
		}
	}

	int_col := FromBase26(colref)
	int_row, _ := strconv.Atoi(rowref)

	return int_row - 1, int_col - 1
}

func CellRefFromCoords(column, row int) string {
	return fmt.Sprintf("%s%d", ToBase26(column+1), row+1)
}
