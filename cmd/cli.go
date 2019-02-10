package cmd

import (
	"os"
	"strconv"
)

func GetRowsAndColumns() (int, int) {
	rows := 5
	columns := 5
	if len(os.Args) == 3 {
		if _rows, err := strconv.Atoi(os.Args[1]); err == nil {
			rows = _rows
		}
		if _columns, err := strconv.Atoi(os.Args[2]); err == nil {
			columns = _columns
		}
	}

	return rows, columns
}
