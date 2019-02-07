package main

import "fmt"

type Grid struct {
	rows, columns int
	cells         [][]*Cell
}

type Cell struct {
	x, y int
}

func main() {
	g := NewGrid(5, 5)
	g.Add(&Cell{1, 2})
	g.Add(&Cell{4, 4})
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", g.At(0, 1))

	g.Each(func(c *Cell) {
		fmt.Printf("cell: %v\n", c)
	})
}

func NewGrid(rows, columns int) Grid {
	g := Grid{rows: rows, columns: columns}
	cells := make([][]*Cell, rows)
	for y := 0; y < rows; y++ {
		cells[y] = make([]*Cell, 0)
		for x := 0; x < columns; x++ {
			cells[y] = append(cells[y], &Cell{x: x, y: y})
		}
	}
	g.cells = cells
	return g
}

func (g *Grid) Add(cell *Cell) {
	g.cells[cell.y][cell.x] = cell
}

func (g *Grid) At(row, column int) *Cell {
	return g.cells[row][column]
}

func (g *Grid) EachXY(fn func(*Cell, int, int)) {
	for y, row := range g.cells {
		for x, cell := range row {
			fn(cell, x, y)
		}
	}
}

func (g *Grid) Each(fn func(*Cell)) {
	g.EachXY(func(c *Cell, _, _ int) {
		fn(c)
	})
}
