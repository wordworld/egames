package game

import (
	"f2v18/board"
	"f2v18/conf"
)

type GridBoard struct {
	Width     int
	Height    int
	Rows      int
	Cols      int
	RowHeight int
	ColWidth  int
	Left      int
	Top       int
}

func NewGridBoard(opt *board.Option) *GridBoard {
	b := &GridBoard{
		Rows: opt.Rows,
		Cols: opt.Cols,
	}
	cfg := conf.GetInstance()
	b.ColWidth = (opt.Width - cfg.WidthFrame*2 - cfg.WidthLine*(b.Cols-2)) / b.Cols
	b.RowHeight = (opt.Height - cfg.WidthFrame*2 - cfg.WidthLine*(b.Rows-2)) / b.Rows
	// 棋盘
	b.Width = cfg.WidthFrame*2 + cfg.WidthLine*(b.Cols-2) + b.ColWidth*(b.Cols-1)
	b.Height = cfg.WidthFrame*2 + cfg.WidthLine*(b.Rows-2) + b.RowHeight*(b.Rows-1)
	b.Left, b.Top = b.ColWidth/2, b.RowHeight/2
	return b
}

// 用行、列定位坐标
func (b *GridBoard) Coord(row, col int) (int, int) {
	cfg := conf.GetInstance()
	x := b.Left + cfg.WidthFrame + (b.ColWidth+cfg.WidthLine)*col
	y := b.Top + cfg.WidthFrame + (b.RowHeight+cfg.WidthLine)*row
	return x, y
}

func (b *GridBoard) GetLeft() int {
	return b.Left
}

func (b *GridBoard) GetRight() int {
	return b.Left + b.Width
}

func (b *GridBoard) GetWidth() int {
	return b.Width
}

func (b *GridBoard) GetHeight() int {
	return b.Height
}

func (b *GridBoard) GetTop() int {
	return b.Top
}

func (b *GridBoard) GetBottom() int {
	return b.Left + b.Height
}
