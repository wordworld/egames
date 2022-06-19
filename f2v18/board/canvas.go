package board

import (
	"f2v18/conf"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

type Canvas struct {
	*Option
	*ebiten.Image
	*ebiten.DrawImageOptions
}

func NewCanvas(options ...Operator) *Canvas {
	v := &Canvas{}
	v.Option = new(Option)
	v.Option.Apply(options...)
	v.Image = ebiten.NewImage(v.Width, v.Height)
	v.DrawImageOptions = &ebiten.DrawImageOptions{}
	return v
}

// 更新游戏对象
func (v *Canvas) UpdateImage() (e *ebiten.Image, o *ebiten.DrawImageOptions) {
	e = v.Image
	o = v.DrawImageOptions
	if !v.UpdateAll {
		return
	}
	// 全量重绘
	defer v.Disposable()
	v.DrawBoard()
	return
}

func (v *Canvas) DrawBoard() {
	cfg := conf.GetInstance()
	v.Fill(cfg.ColorBoard)
	colWidth := (v.Width-2*cfg.WidthFrame-(v.Cols-2)*cfg.WidthLine)/v.Cols + 1
	rowHeight := (v.Height-2*cfg.WidthFrame-(v.Rows-2)*cfg.WidthLine)/v.Rows + 1
	// 棋盘
	width, height := v.Width-colWidth, v.Height-rowHeight
	left, top := colWidth/2, rowHeight/2
	// outter-lines
	right, bottom := v.Width-left-cfg.WidthFrame, v.Height-top-cfg.WidthFrame
	v.DrawRect(left, top, width, cfg.WidthFrame, cfg.ColorLine)
	v.DrawRect(left, bottom, width, cfg.WidthFrame, cfg.ColorLine)
	v.DrawRect(left, top, cfg.WidthFrame, height, cfg.ColorLine)
	v.DrawRect(right, top, cfg.WidthFrame, height, cfg.ColorLine)
	// inner-horizon-lines
	for y := top + cfg.WidthFrame + rowHeight; y < bottom; y += cfg.WidthLine + rowHeight {
		v.DrawRect(left, y, width, cfg.WidthLine, cfg.ColorLine)
	}
	// inner-vertical-lines
	for x := left + cfg.WidthFrame + colWidth; x < right; x += cfg.WidthLine + colWidth {
		v.DrawRect(x, top, cfg.WidthLine, height, cfg.ColorLine)
	}
}

func (v *Canvas) DrawRect(left, top, width, height int, color color.Color) {
	ebitenutil.DrawRect(v.Image, float64(left), float64(top), float64(width), float64(height), color)
}

func (v *Canvas) DrawLine() {

}
