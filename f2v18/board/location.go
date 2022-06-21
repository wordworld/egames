package board

type Location interface {
	Coord(row, col int) (int, int)
	GetLeft() int
	GetTop() int
	GetWidth() int
	GetHeight() int
	GetRight() int
	GetBottom() int
}
