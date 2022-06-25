package view

import (
	"f2v18/board"
	"github.com/hajimehoshi/ebiten/v2"
)

type ViewManager struct {
	*board.Canvas
}

func (cvs *ViewManager) DrawDynamic() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return nil, nil
}
