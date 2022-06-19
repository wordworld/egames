package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Game struct {
	img *ebiten.Image
}

func NewGame() (g *Game) {
	g = new(Game)
	file := "algorithm.png"
	reader, err := os.Open(file)
	defer reader.Close()
	if err != nil {
		log.Fatalf("read file failed:%v", err)
	}
	img, err := png.Decode(reader)
	if err != nil {
		log.Fatalf("image decode failed:%v", err)
	}
	// ebitenutil.NewImageFromFile 对一些图片格式不支持
	// 用 image.Image 创建 ebiten.Image 比较保险
	g.img = ebiten.NewImageFromImage(img)
	if g.img == nil {
		log.Fatalf("new image failed")
	}
	return
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w, h := screen.Size()
	screen.Fill(color.RGBA{128, 0, 0, 255})
	ebitenutil.DebugPrintAt(screen, "Hello, World!", w/2, 0)

	opt := &ebiten.DrawImageOptions{}
	sx, sy := 0.5, 0.5
	opt.GeoM.Scale(sx, sy)

	iw, ih := g.img.Size()
	dx := float64(iw) * sx / 2
	dy := float64(ih) * sy / 2
	opt.GeoM.Translate(float64(w)/2-dx, float64(h)/2-dy)
	screen.DrawImage(g.img, opt)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 300
}
