package main

import "fmt"

type Grid struct {
	rows, columns int
	cells         [][]*Cell
}

type Cell struct {
	x, y                     int
	north, east, south, west *Cell
}

func main() {
	g := NewGrid(5, 5)
	g.Add(&Cell{y: 1, x: 2})
	g.Add(&Cell{y: 4, x: 4})
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", g.At(0, 1))

	g.Each(func(c *Cell) {
		fmt.Printf("cell: %v\n", c)
	})

	fmt.Printf("north of (0,1): %v\n", g.At(0, 1).north)
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
	g.setNeighbors()
	return g
}

func (g *Grid) Add(cell *Cell) {
	g.cells[cell.y][cell.x] = cell
}

func (g *Grid) At(x, y int) *Cell {
	if y < 0 || x < 0 || y >= len(g.cells) || x >= len(g.cells[y]) {
		return nil
	} else {
		return g.cells[y][x]
	}
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

func (g *Grid) setNeighbors() {
	g.Each(func(cell *Cell) {
		cell.north = g.At(cell.y-1, cell.x)
		cell.south = g.At(cell.y+1, cell.x)
		cell.east = g.At(cell.y, cell.x+1)
		cell.west = g.At(cell.y, cell.x-1)
	})
}
