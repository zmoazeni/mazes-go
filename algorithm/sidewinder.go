package algorithm

import "math/rand"

func Sidewinder(grid *Grid) {
	grid.EachRow(func(row []*Cell) {
		run := make([]*Cell, 0)

		for _, cell := range row {
			run = append(run, cell)
			atEasternBorder := cell.East == nil
			atSouthernBorder := cell.South == nil
			shouldClose := atEasternBorder || (!atSouthernBorder && rand.Intn(2) == 0)

			if shouldClose {
				member := run[rand.Intn(len(run))]
				if member.North != nil {
					member.Link(member.North)
				}
			} else {
				cell.Link(cell.East)
			}
		}
	})
}
