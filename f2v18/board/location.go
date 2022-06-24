package board

import "f2v18/template/gen/tint"

type Location interface {
	Coord(row, col int) (int, int)
	Quad(row, col int) *tint.Quad
	GetLeft() int
	GetTop() int
	GetWidth() int
	GetHeight() int
	GetRight() int
	GetBottom() int
}
