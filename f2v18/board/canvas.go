package board

import (
	"f2v18/conf"
	"f2v18/util"
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
	Location
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
		"Radius":      float32(r - 1),
		"Center":      []float32{float32(r), float32(r)},
		"Antialias":   conf.GetInstance().Antialias,
		"Color":       util.Color2Vec(color),
		"ShadowColor": util.Color2Vec(util.InvColor(color)),
		"ShadowDist":  float32(r-1) * 3,
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
	// inner-horizon-lines
	v.DrawRect(v.GetLeft(), v.GetTop(), v.GetWidth(), cfg.WidthFrame, cfg.ColorLine)
	for row := 1; row < v.Rows-1; row++ {
		x, y := v.Coord(row, 0)
		v.DrawRect(x, y, v.GetWidth(), cfg.WidthLine, cfg.ColorLine)
	}
	v.DrawRect(v.GetLeft(), v.GetBottom(), v.GetWidth()+cfg.WidthFrame, cfg.WidthFrame, cfg.ColorLine)
	// inner-vertical-lines
	v.DrawRect(v.GetLeft(), v.GetTop(), cfg.WidthFrame, v.GetHeight(), cfg.ColorLine)
	for col := 1; col < v.Cols-1; col++ {
		x, y := v.Coord(0, col)
		v.DrawRect(x, y, cfg.WidthLine, v.GetHeight(), cfg.ColorLine)
	}
	v.DrawRect(v.GetRight(), v.GetTop(), cfg.WidthFrame, v.GetHeight(), cfg.ColorLine)
}

func (v *Canvas) DrawRect(left, top, width, height int, color color.Color) {
	ebitenutil.DrawRect(v.Image, float64(left), float64(top), float64(width), float64(height), color)
}

func (v *Canvas) DrawPieces() {
	r := conf.GetInstance().RadiusPiece
	for row := 0; row < 3; row++ {
		for col := 0; col < v.Cols; col++ {
			x, y := v.Coord(row, col)

			v.DrawPiece(0, x-r, y-r)
		}
	}
	x2, y2 := v.Coord(v.Rows-1, v.Cols/2-1)
	x1, y1 := v.Coord(v.Rows-1, v.Cols/2)
	v.DrawPiece(1, x1-r, y1-r)
	v.DrawPiece(1, x2-r, y2-r)
}

func (v *Canvas) DrawPiece(i, x, y int) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(x), float64(y))
	v.DrawImage(v.Piece[i], opt)
}
