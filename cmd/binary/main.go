package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	. "github.com/zmoazeni/mazes-go/algorithm"
	. "github.com/zmoazeni/mazes-go/formatter"
)

func main() {
	rand.Seed(time.Now().UnixNano())

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
	grid := NewGrid(rows, columns)
	BinaryTree(&grid)
	fmt.Print(AsciiFormatter(&grid))
	PNG(&grid, 150) // nolint
}
