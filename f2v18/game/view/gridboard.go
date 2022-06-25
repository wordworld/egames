package view

import (
	"f2v18/board"
	"f2v18/conf"
	"f2v18/template/gen/tint"
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

func NewGridBoard(scrn *board.Option) *GridBoard {
	b := &GridBoard{
		Rows: scrn.Rows,
		Cols: scrn.Cols,
	}
	cfg := conf.GetInstance()
	// 整个 screen 由 Rows 行 Cols 列，均匀划分成规则网格
	// ColWidth 是相邻两条竖线间空白区的宽度，屏幕左右各有 ColWidth/2 空白
	b.ColWidth = (scrn.Width - cfg.WidthFrame*2 - cfg.WidthLine*(b.Cols-2)) / b.Cols
	// RowHeight 是相邻两条横线间空白的高度，屏幕上下各有 RowHeight/2 空白
	b.RowHeight = (scrn.Height - cfg.WidthFrame*2 - cfg.WidthLine*(b.Rows-2)) / b.Rows
	// 棋盘宽度：左框左沿 ～ 右框右沿
	b.Width = cfg.WidthFrame*2 + cfg.WidthLine*(b.Cols-2) + b.ColWidth*(b.Cols-1)
	// 棋盘高度：上框上沿 ～ 下框下沿
	b.Height = cfg.WidthFrame*2 + cfg.WidthLine*(b.Rows-2) + b.RowHeight*(b.Rows-1)
	// 左上角
	b.Left, b.Top = (scrn.Width-b.Width)/2, (scrn.Height-b.Height)/2
	return b
}

func (b *GridBoard) Coord(row, col int) (int, int) {
	q := b.Quad(row, col)
	return *q.X, *q.Y
}

// 用行、列定位坐标：第row行、第col列相交矩形中心(x,y) 及相交矩形的宽u、高v
func (b *GridBoard) Quad(row, col int) *tint.Quad {
	cfg := conf.GetInstance()
	x, y := b.Left, b.Top
	u, v := cfg.WidthLine, cfg.WidthLine
	if col == 0 {
		x += cfg.WidthFrame / 2
		u = cfg.WidthFrame
	} else {
		x += cfg.WidthFrame + b.ColWidth + cfg.WidthLine/2
		if col == b.Cols-1 {
			x = b.Left + b.Width - cfg.WidthFrame/2
			u = cfg.WidthFrame
		} else {
			x += (b.ColWidth + cfg.WidthLine) * (col - 1)
		}
	}
	if row == 0 {
		y += cfg.WidthFrame / 2
		v = cfg.WidthFrame
	} else {
		y += cfg.WidthFrame + b.RowHeight + cfg.WidthLine/2
		if row == b.Rows-1 {
			y = b.Top + b.Height - cfg.WidthFrame/2
			v = cfg.WidthFrame
		} else {
			y += (b.RowHeight + cfg.WidthLine) * (row - 1)
		}
	}
	q := &tint.Quad{}
	q.Set(x, y, u, v)
	return q
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
