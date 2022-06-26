package game

import (
	"f2v18/board"
	"f2v18/conf"
	"f2v18/game/input"
	"f2v18/game/view"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	CONFIG_FILE = "game.conf.json"
)

type Game struct {
	*view.ViewManager // 屏幕显示
	*input.Stroke     // 鼠标/触控
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
	var justStroke *input.Stroke
	//鼠标单击
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		justStroke = input.NewStroke(&input.MouseStrokeSource{})
	}
	// 按下触控
	touchIDs := make([]ebiten.TouchID, 0)
	touchIDs = inpututil.AppendJustPressedTouchIDs(touchIDs)
	if touched := len(touchIDs); touched > 0 {
		justStroke = input.NewStroke(&input.TouchStrokeSource{touchIDs[touched-1]})
	}
	// 处理之前的操作
	g.UpdateStroke()
	// 替换处理最新的
	g.onNewStroke(justStroke)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.ViewManager.UpdateImage())
	g.DrawDynamic(screen)
	if g.Stroke != nil {
		x, y := g.Stroke.Position()
		row, col := g.LocateCoord(x, y)
		index := g.GetIndex(row, col)
		ebitenutil.DebugPrint(screen, fmt.Sprintf("clicked %v(%v,%v) nearest grid %v(%v,%v)", g.SelIndex, x, y, index, row, col))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Canvas.Size()
}

func (g *Game) Quit() {
	conf.GetInstance().Save(CONFIG_FILE)
}

func (g *Game) UpdateStroke() {
	if g.Stroke == nil {
		return
	}
	// 更新位置状态
	g.Stroke.Update()
	if g.Stroke.IsReleased() { // 释放
		g.ReleaseSelect(g.Stroke.Position())
		g.Stroke = nil
	} else {
		g.UpdateSelect(g.Stroke.Position())
	}

}

func (g *Game) onNewStroke(stroke *input.Stroke) {
	if stroke == nil {
		return
	}
	if g.Stroke != nil { // 还没释放
		return
	}
	g.Stroke = stroke
	g.Stroke.Update()
	g.SetSelect(g.Stroke.Position())
}
