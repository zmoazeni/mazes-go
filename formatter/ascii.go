package formatter

import (
	"strings"

	a "github.com/zmoazeni/mazes-go/algorithm"
)

func AsciiFormatter(grid *a.Grid) string {
	var out strings.Builder
	out.WriteString("+")
	out.WriteString(strings.Repeat("---+", grid.Columns))
	out.WriteString("\n")

	body := strings.Repeat(" ", 3)
	corner := "+"

	grid.EachRow(func(row []*a.Cell) {
		var top, bottom strings.Builder
		top.WriteString("|")
		bottom.WriteString("+")

		for _, cell := range row {
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
