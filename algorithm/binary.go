package algorithm

import "math/rand"

func BinaryTree(grid *Grid) {
	grid.Each(func(cell *Cell) {
		neighbors := make([]*Cell, 0)
		if cell.North != nil {
			neighbors = append(neighbors, cell.North)
		}

		if cell.East != nil {
			neighbors = append(neighbors, cell.East)
		}

		if len(neighbors) > 0 {
			index := rand.Intn(len(neighbors))
			neighbor := neighbors[index]
			cell.Link(neighbor)
		}
	})
}
