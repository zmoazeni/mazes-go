package algorithm

type Grid struct {
	Rows, Columns int
	cells         [][]*Cell
}

type Cell struct {
	x, y                     int
	North, East, South, West *Cell
	links                    map[*Cell]bool
}

func NewCell(x, y int) Cell {
	cell := Cell{x: x, y: y}
	cell.links = make(map[*Cell]bool)
	return cell
}

func NewGrid(rows, columns int) Grid {
	g := Grid{Rows: rows, Columns: columns}
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

func (g *Grid) EachRow(fn func([]*Cell)) {
	for _, row := range g.cells {
		fn(row)
	}
}

func (g *Grid) setNeighbors() {
	g.Each(func(cell *Cell) {
		cell.North = g.At(cell.x, cell.y-1)
		cell.South = g.At(cell.x, cell.y+1)
		cell.East = g.At(cell.x+1, cell.y)
		cell.West = g.At(cell.x-1, cell.y)
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
	if c.North != nil {
		neighbors = append(neighbors, c.North)
	}

	if c.South != nil {
		neighbors = append(neighbors, c.South)
	}

	if c.East != nil {
		neighbors = append(neighbors, c.East)
	}

	if c.West != nil {
		neighbors = append(neighbors, c.West)
	}

	return neighbors
}
