package tokenizer

import "fmt"

func InitTable(table *[MAX_ROWS][MAX_COLUMNS]int) {
	for i := 0; i < MAX_ROWS; i++ {
		for j := 0; j < MAX_COLUMNS; j++ {
			(*table)[i][j] = -1
		}
	}
}

func MakeSuccess(table *[MAX_ROWS][MAX_COLUMNS]int, state int) {
	if state < 0 || state >= MAX_ROWS {
		panic("Invalid state")
	}

	(*table)[state][0] = 1
}

func MakeFail(table *[MAX_ROWS][MAX_COLUMNS]int, state int) {
	if state < 0 || state >= MAX_ROWS {
		panic("Invalid state")
	}

	(*table)[state][0] = 0
}

func IsSuccess(table *[MAX_ROWS][MAX_COLUMNS]int, state int) bool {
	if state < 0 || state >= MAX_ROWS {
		panic("Invalid state")
	}

	return (*table)[state][0] == 1
}

func MarkCell(row int, table *[MAX_ROWS][MAX_COLUMNS]int, column int, state int) {
	if row < 0 || row >= MAX_ROWS {
		panic("Invalid row")
	}
	if state < 0 || state >= MAX_ROWS {
		panic("Invalid state")
	}
	if column < 0 || column+1 >= MAX_COLUMNS {
		panic("Invalid column")
	}

	(*table)[row][column+1] = state
}

func MarkCells(
	row int,
	table *[MAX_ROWS][MAX_COLUMNS]int,
	from int,
	to int,
	state int,
) {
	if row < 0 || row >= MAX_ROWS {
		panic("Invalid row")
	}
	if state < 0 || state >= MAX_ROWS {
		panic("Invalid state")
	}

	for i := from + 1; i <= to+1; i++ { // Set value to the range of columns
		if i < 0 || i+1 >= MAX_COLUMNS {
			panic("Invalid column")
		}
		(*table)[row][i] = state
	}
}

func MarkChars(
	row int,
	table *[MAX_ROWS][MAX_COLUMNS]int,
	s string,
	state int,
) {
	if row < 0 || row >= MAX_ROWS {
		panic("Invalid row")
	}
	if state < 0 || state >= MAX_ROWS {
		panic("Invalid state")
	}

	for i := 0; i < len(s); i++ {
		if s[i]+1 >= byte(MAX_COLUMNS) {
			panic("Invalid column")
		}

		(*table)[row][s[i]+1] = state
	}
}

func PrintTable(table *[MAX_ROWS][MAX_COLUMNS]int) {
	for i := 0; i < MAX_ROWS; i++ {
		for j := 0; j < MAX_COLUMNS; j++ {
			fmt.Printf("[%3d] ", (*table)[i][j])
		}
		fmt.Println()
	}
}

func ShowStr(s string, pos int) {
	fmt.Printf("%s [pos: %d]\n", s, pos)
	for i := 0; i < len(s); i++ {
		if i == pos {
			fmt.Printf("^")
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Println()
}
