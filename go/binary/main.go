package main

import "fmt"

type Grid struct {
	rows, columns int
	cells         [][]*Cell
}

type Cell struct {
	row, column int
}

type GridIter struct {
	fetched                   bool
	currentRow, currentColumn int
	grid                      *Grid
}

func main() {
	g := NewGrid(5, 5)
	g.Add(&Cell{1, 2})
	g.Add(&Cell{4, 4})
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", g.At(0, 1))

	iter := g.Iter()
	for iter.Next() {
		c := iter.Get()
		fmt.Printf("Cell: %v\n", c)
	}
}

func NewGrid(rows, columns int) Grid {
	g := Grid{rows: rows, columns: columns}
	cells := make([][]*Cell, rows)
	for i := 0; i < rows; i++ {
		cells[i] = make([]*Cell, columns)
	}
	g.cells = cells
	return g
}

func (g *Grid) Add(cell *Cell) {
	g.cells[cell.row][cell.column] = cell
}

func (g *Grid) At(row, column int) *Cell {
	return g.cells[row][column]
}

func (g *Grid) Iter() GridIter {
	return GridIter{grid: g}
}

func (iter *GridIter) Next() bool {
	// Leave pointer at 0,0 for the first iteration
	if !iter.fetched {
		iter.fetched = true
		return iter.grid.rows > 0 && iter.grid.columns > 0
	}

	if iter.currentColumn < iter.grid.columns-1 {
		iter.currentColumn++
		return true
	} else if iter.currentRow < iter.grid.rows-1 {
		iter.currentRow++
		iter.currentColumn = 0
		return true
	} else {
		return false
	}
}

func (iter *GridIter) Get() *Cell {
	return iter.grid.At(iter.currentRow, iter.currentColumn)
}
