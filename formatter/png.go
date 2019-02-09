package formatter

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
	a "github.com/zmoazeni/mazes-go/algorithm"
)

func drawBackground(gc *draw2dimg.GraphicContext, r image.Rectangle) {
	gc.SetFillColor(color.White)
	gc.MoveTo(float64(r.Min.X), float64(r.Min.Y))
	gc.LineTo(float64(r.Max.X), float64(r.Min.Y))
	gc.LineTo(float64(r.Max.X), float64(r.Max.Y))
	gc.LineTo(float64(r.Min.X), float64(r.Max.Y))
	gc.Close()
	gc.FillStroke()
}

func PNG(grid *a.Grid, cellSize int) error {
	width := cellSize * grid.Columns
	height := cellSize * grid.Rows

	rect := image.Rect(0, 0, width, height)
	dest := image.NewRGBA(rect)
	gc := draw2dimg.NewGraphicContext(dest)
	drawBackground(gc, rect)

	grid.Each(func(cell *a.Cell) {
		gc.BeginPath()
		gc.SetLineWidth(5)
		gc.SetStrokeColor(color.Black)
		x1 := float64(cell.X * cellSize)
		y1 := float64(cell.Y * cellSize)
		x2 := float64((cell.X + 1) * cellSize)
		y2 := float64((cell.Y + 1) * cellSize)

		if cell.North == nil || !cell.IsLinked(cell.North) {
			gc.MoveTo(x1, y1)
			gc.LineTo(x2, y1)
		}

		if cell.West == nil || !cell.IsLinked(cell.West) {
			gc.MoveTo(x1, y1)
			gc.LineTo(x1, y2)
		}
		gc.FillStroke()
	})

	return draw2dimg.SaveToPngFile("matrix.png", dest)
}
