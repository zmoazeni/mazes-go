package formatter

import (
	"fmt"
	"strings"

	a "github.com/zmoazeni/mazes-go/algorithm"
)

func AsciiFormatterWithDistance(grid *a.Grid, distances *a.Distances) string {
	var out strings.Builder
	out.WriteString("+")
	out.WriteString(strings.Repeat("---+", grid.Columns))
	out.WriteString("\n")

	corner := "+"

	grid.EachRow(func(row []*a.Cell) {
		var top, bottom strings.Builder
		top.WriteString("|")
		bottom.WriteString("+")
		defaultBody := strings.Repeat(" ", 3)

		for _, cell := range row {
			var body string
			if distances != nil {
				d, found := distances.DistanceFromRoot(cell)
				if found {
					body = fmt.Sprintf("%03d", d)
				} else {
					body = defaultBody
				}
			} else {
				body = defaultBody
			}

			eastBoundary := "|"
			if cell.IsLinked(cell.East) {
				eastBoundary = " "
			}
			top.WriteString(body)
			top.WriteString(eastBoundary)

			southBoundary := "---"
			if cell.IsLinked(cell.South) {
				southBoundary = strings.Repeat(" ", 3)
			}

			bottom.WriteString(southBoundary)
			bottom.WriteString(corner)
		}

		out.WriteString(top.String())
		out.WriteString("\n")
		out.WriteString(bottom.String())
		out.WriteString("\n")
	})

	return out.String()
}

func AsciiFormatter(grid *a.Grid) string {
	return AsciiFormatterWithDistance(grid, nil)
}
