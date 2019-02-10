package main

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/zmoazeni/mazes-go/algorithm"
	. "github.com/zmoazeni/mazes-go/cmd"
	. "github.com/zmoazeni/mazes-go/formatter"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	rows, columns := GetRowsAndColumns()
	grid := NewGrid(rows, columns)
	Sidewinder(&grid)
	fmt.Print(AsciiFormatter(&grid))
	PNG(&grid, 150, "sidewinder.png") // nolint
}