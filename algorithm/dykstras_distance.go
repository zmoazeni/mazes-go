package algorithm

type Distances struct {
	root      *Cell
	distances map[*Cell]int
}

func newDistances(cell *Cell) Distances {
	distances := make(map[*Cell]int)
	distances[cell] = 0
	return Distances{
		root:      cell,
		distances: distances,
	}
}

func (d *Distances) DistanceFromRoot(cell *Cell) (int, bool) {
	val, ok := d.distances[cell]
	return val, ok
}

func CalculateDistances(grid *Grid, x, y int) Distances {
	rootCell := grid.At(x, y)
	d := newDistances(rootCell)
	frontier := []*Cell{rootCell}

	for len(frontier) > 0 {
		newFrontier := make([]*Cell, 0)
		for _, cell := range frontier {
			for _, linkedCell := range cell.Links() {
				if _, ok := d.distances[linkedCell]; !ok {
					d.distances[linkedCell] = d.distances[cell] + 1
					newFrontier = append(newFrontier, linkedCell)
				}
			}
		}
		frontier = newFrontier
	}

	return d
}
