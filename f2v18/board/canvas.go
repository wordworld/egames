package board

import (
	"f2v18/conf"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

type Canvas struct {
	*Option
	*ebiten.Image
	*ebiten.DrawImageOptions
	Piece [2]*ebiten.Image
}

func NewCanvas(options ...Operator) *Canvas {
	v := &Canvas{}
	v.Option = new(Option)
	v.Option.Apply(options...)
	v.Image = ebiten.NewImage(v.Width, v.Height)
	v.DrawImageOptions = &ebiten.DrawImageOptions{}
	for i, c := range conf.GetInstance().ColorPieces {
		v.Piece[i] = NewPiece(c)
	}
	return v
}

func NewPiece(color *color.RGBA) *ebiten.Image {
	shader, err := ebiten.NewShader(shader_pieces)
	if err != nil {
		log.Fatalf("piece shader create err:%v\n", err)
		return nil
	}
	r := conf.GetInstance().RadiusPiece
	opt := &ebiten.DrawRectShaderOptions{}
	opt.Uniforms = map[string]interface{}{
		"Radius": float32(r - 1),
		"Center": []float32{float32(r), float32(r)},
		"Color":  []float32{float32(color.R) / 255, float32(color.G) / 255, float32(color.B) / 255, float32(color.A) / 255},
	}
	w, h := 2*r, 2*r
	piece := ebiten.NewImage(w, h)
	piece.DrawRectShader(w, h, shader, opt)
	return piece
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
	v.DrawPieces()
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

func (v *Canvas) DrawPieces() {
	v.DrawPiece(0, 20, 50)
	v.DrawPiece(1, 100, 50)
}

func (v *Canvas) DrawPiece(i, x, y int) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(x), float64(y))
	v.DrawImage(v.Piece[i], opt)
}
