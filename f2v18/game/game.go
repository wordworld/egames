package game

import (
	"f2v18/board"
	"f2v18/conf"
	"f2v18/game/view"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CONFIG_FILE = "game.conf.json"
)

type Game struct {
	*view.ViewManager
}

func NewGame() (g *Game) {
	cfg, _ := conf.GetInstance().Load(CONFIG_FILE)
	scrnWidth, scrnHeight := ebiten.ScreenSizeInFullscreen()
	ebiten.SetWindowSize(scrnWidth-cfg.MarginHor, scrnHeight-cfg.MarginVer)
	ebiten.SetWindowTitle(cfg.GameName)
	g = &Game{ViewManager: &view.ViewManager{}}
	g.Canvas = board.NewCanvas(board.WithSize(ebiten.WindowSize()),
		board.WithGrid(cfg.LnHorizon, cfg.LnVertical),
		board.WithLocation(view.NewGridBoard),
		board.WidthUpdateAll())
	return
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Canvas.UpdateImage())
	g.ViewManager.DrawPieces()
	if dyn, opt := g.DrawDynamic(); dyn != nil {
		screen.DrawImage(dyn, opt)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Canvas.Size()
}

func (g *Game) Quit() {
	conf.GetInstance().Save(CONFIG_FILE)
}
