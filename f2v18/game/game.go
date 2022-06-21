package game

import (
	"f2v18/board"
	"f2v18/conf"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CONFIG_FILE = "game.conf.json"
)

type Game struct {
	canvas *board.Canvas
}

func NewGame() (g *Game) {
	cfg, _ := conf.GetInstance().Load(CONFIG_FILE)
	g = new(Game)
	ebiten.SetWindowTitle(cfg.GameName)
	ebiten.SetWindowSize(cfg.WinWidth, cfg.WinHeight)
	rows, cols := 6, 6
	g.canvas = board.NewCanvas(board.WithSize(ebiten.WindowSize()),
		board.WithGrid(rows, cols),
		board.WidthUpdateAll())
	g.canvas.Location = NewGridBoard(g.canvas.Option)
	return
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.canvas.UpdateImage())
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.canvas.Size()
}

func (g *Game) Quit() {
	conf.GetInstance().Save(CONFIG_FILE)
}
