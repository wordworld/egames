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
	*Photo
	Piece [2]*Photo
	Location
}

func NewCanvas(options ...Operator) *Canvas {
	opt := new(Option)
	opt.Apply(options...)
	cvs := &Canvas{
		Option:   opt,
		Photo:    NewPhoto(opt.Width, opt.Height),
		Piece:    [2]*Photo{},
		Location: nil,
	}
	for i, c := range conf.GetInstance().ColorPieces {
		cvs.Piece[i] = TakePhoto(NewPiece(c))
	}
	return cvs
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
		"Radius":      float32(r),
		"Center":      []float32{float32(r), float32(r)},
		"Antialias":   conf.GetInstance().Antialias,
		"Color":       util.Color2Vec(color),
		"ShadowColor": util.Color2Vec(util.InvColor(color)),
		"ShadowDist":  float32(r) * 3.2,
		"HighLight":   float32(2),
	}
	w, h := 2*r, 2*r
	piece := ebiten.NewImage(w, h)
	piece.DrawRectShader(w, h, shader, opt)
	return piece
}

// 更新游戏对象
func (cvs *Canvas) UpdateImage() (e *ebiten.Image, o *ebiten.DrawImageOptions) {
	e = cvs.Image
	o = cvs.DrawImageOptions
	if !cvs.UpdateAll {
		return
	}
	// 全量重绘
	defer cvs.Disposable()
	cvs.DrawBoard()
	return
}

func (cvs *Canvas) DrawBoard() {
	cfg := conf.GetInstance()
	cvs.Fill(cfg.ColorBoard)
	// horizon-lines
	for row := 0; row < cvs.Rows; row++ {
		x, y, u, v := cvs.Quad(row, 0).QUAD()
		cvs.DrawRect(x-u/2, y-v/2, cvs.GetWidth(), v, cfg.ColorLine)
	}
	// vertical-lines
	for col := 0; col < cvs.Cols; col++ {
		x, y, u, v := cvs.Quad(0, col).QUAD()
		cvs.DrawRect(x-u/2, y-v/2, u, cvs.GetHeight(), cfg.ColorLine)
	}
}

func (cvs *Canvas) DrawRect(left, top, width, height int, color color.Color) {
	ebitenutil.DrawRect(cvs.Image, float64(left), float64(top), float64(width), float64(height), color)
}

func (cvs *Canvas) DrawPiece(x, y int, side int) {
	cvs.Piece[side].Take().Put(x, y).Print(cvs.Image)
}

// 落子
func (cvs *Canvas) PutPiece(row, col int, side int) {
	x, y := cvs.Coord(row, col)
	r := conf.GetInstance().RadiusPiece
	cvs.DrawPiece(x-r, y-r, side)
}

// 取子
func (cvs *Canvas) TakePiece(row, col int) {
	x, y, u, v := cvs.Quad(row, col).QUAD()
	cfg := conf.GetInstance()
	r := cfg.RadiusPiece
	cvs.DrawRect(x-r, y-r, 2*r, 2*r, cfg.ColorBoard)
	// 竖线
	if 0 != row {
		cvs.DrawRect(x-u/2, y-r, u, r, cfg.ColorLine) // up
	}
	if row != cvs.Rows-1 {
		cvs.DrawRect(x-u/2, y, u, r, cfg.ColorLine) // down
	}
	// 横线
	if 0 != col {
		cvs.DrawRect(x-u/2-r, y-v/2, r+u, v, cfg.ColorLine) // left
	}
	if col != cvs.Cols-1 {
		cvs.DrawRect(x-u/2, y-v/2, r+u, v, cfg.ColorLine) // right
	}
}
