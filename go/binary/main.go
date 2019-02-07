package main

import (
	"fmt"
	"math/rand"
)

type Grid struct {
	rows, columns int
	cells         [][]*Cell
}

type Cell struct {
	x, y                     int
	north, east, south, west *Cell
	links                    map[*Cell]bool
}

func main() {
	grid := NewGrid(5, 5)
	BinaryTree(&grid)
	fmt.Printf("%v\n", grid)
}

func NewCell(x, y int) Cell {
	cell := Cell{x: x, y: y}
	cell.links = make(map[*Cell]bool)
	return cell
}

func NewGrid(rows, columns int) Grid {
	g := Grid{rows: rows, columns: columns}
	cells := make([][]*Cell, rows)
	for y := 0; y < rows; y++ {
		cells[y] = make([]*Cell, 0, columns)
		for x := 0; x < columns; x++ {
			cell := NewCell(x, y)
			cells[y] = append(cells[y], &cell)
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
		cell.north = g.At(cell.x, cell.y-1)
		cell.south = g.At(cell.x, cell.y+1)
		cell.east = g.At(cell.x+1, cell.y)
		cell.west = g.At(cell.x-1, cell.y)
	})
}

func (c *Cell) Link(otherCell *Cell) {
	c.LinkSingle(otherCell)
	otherCell.LinkSingle(c)
}

func (c *Cell) LinkSingle(otherCell *Cell) {
	c.links[otherCell] = true
}

func (c *Cell) Unlink(otherCell *Cell) {
	c.UnlinkSingle(otherCell)
	otherCell.UnlinkSingle(c)
}

func (c *Cell) UnlinkSingle(otherCell *Cell) {
	delete(c.links, otherCell)
}

func (c *Cell) Links() []*Cell {
	keys := make([]*Cell, 0, len(c.links))
	for cell := range c.links {
		keys = append(keys, cell)
	}
	return keys
}

func (c *Cell) IsLinked(otherCell *Cell) bool {
	_, found := c.links[otherCell]
	return found
}

func (c *Cell) Neighbors() []*Cell {
	neighbors := make([]*Cell, 0)
	if c.north != nil {
		neighbors = append(neighbors, c.north)
	}

	if c.south != nil {
		neighbors = append(neighbors, c.south)
	}

	if c.east != nil {
		neighbors = append(neighbors, c.east)
	}

	if c.west != nil {
		neighbors = append(neighbors, c.west)
	}

	return neighbors
}

func BinaryTree(grid *Grid) {
	grid.Each(func(cell *Cell) {
		neighbors := make([]*Cell, 0)
		if cell.north != nil {
			neighbors = append(neighbors, cell.north)
		}

		if cell.east != nil {
			neighbors = append(neighbors, cell.east)
		}

		if len(neighbors) > 0 {
			index := rand.Intn(len(neighbors))
			neighbor := neighbors[index]
			cell.Link(neighbor)
		}
	})
}
