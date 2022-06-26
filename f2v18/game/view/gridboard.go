package view

import (
	"f2v18/board"
	"f2v18/conf"
	"f2v18/template/gen/tint"
	"math"
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
	CY        []int // 横线中心 y 值
	CX        []int // 竖线中心 x 值
}

func NewGridBoard(scrn *board.Option) board.Location {
	b := &GridBoard{
		Rows: scrn.Rows,
		Cols: scrn.Cols,
		CY:   make([]int, scrn.Rows),
		CX:   make([]int, scrn.Cols),
	}
	cfg := conf.GetInstance()
	b.FitScreen(scrn)
	// 调整线宽
	if b.ColWidth < 2*cfg.RadiusPiece || b.RowHeight < 2*cfg.RadiusPiece {
		cfg.WidthFrame = 2
		cfg.WidthLine = 1
		b.FitScreen(scrn)
	}
	// 调整棋子半径
	fitR := int(math.Min(float64(b.ColWidth), float64(b.RowHeight)) * 0.5)
	if cfg.RadiusPiece > fitR {
		cfg.RadiusPiece = fitR
		if cfg.RadiusPiece < 1 {
			cfg.RadiusPiece = 1
		}
	}
	return b
}
func (b *GridBoard) FitScreen(scrn *board.Option) {
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
	// 横线中心
	b.CY[0] = b.Top + cfg.WidthFrame/2
	b.CY[b.Rows-1] = b.GetBottom() - cfg.WidthFrame/2
	for i := 1; i < b.Rows-1; i++ {
		b.CY[i] = b.Top + cfg.WidthFrame + i*(b.RowHeight+cfg.WidthLine) - cfg.WidthLine/2
	}
	// 竖线中心
	b.CX[0] = b.Left + cfg.WidthFrame/2
	b.CX[b.Cols-1] = b.GetRight() - cfg.WidthFrame/2
	for j := 1; j < b.Cols-1; j++ {
		b.CX[j] = b.Left + cfg.WidthFrame + j*(b.ColWidth+cfg.WidthLine) - cfg.WidthLine/2
	}
}

// x -> col  &&  y -> row
func (b *GridBoard) LocateCoord(x, y int) (row int, col int) {
	row, col = tint.LowerBound(b.CY, y), tint.LowerBound(b.CX, x)
	xc, _ := b.GetCoord(0, col)
	if x+b.ColWidth/2 < xc {
		col -= 1
	}
	_, yc := b.GetCoord(row, 0)
	if y+b.RowHeight/2 < yc {
		row -= 1
	}
	return
}

func (b *GridBoard) LocateIndex(index int) (int, int) {
	return index / b.Cols, index % b.Rows
}

func (b *GridBoard) GetCoord(row, col int) (int, int) {
	if row < 0 || b.Rows <= row || col < 0 || b.Cols <= col {
		return 0, 0
	}
	return b.CX[col], b.CY[row]
}

func (b *GridBoard) GetIndex(row, col int) int {
	return row*b.Cols + col
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
	return b.Top + b.Height
}
