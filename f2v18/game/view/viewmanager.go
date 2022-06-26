package view

import (
	"f2v18/board"
	"github.com/hajimehoshi/ebiten/v2"
)

type ViewManager struct {
	*board.Canvas
	SelIndex  int
	PieceSel  *board.Photo // 选中的棋子
	Piece2Put *board.Photo // 即将落子位置标识
}

// 更新游戏对象
func (cvs *ViewManager) UpdateImage() (e *ebiten.Image, o *ebiten.DrawImageOptions) {
	e = cvs.Image
	o = cvs.DrawImageOptions
	if !cvs.UpdateAll {
		return
	}
	// 全量重绘
	defer cvs.Disposable()
	cvs.DrawBoard()
	cvs.DrawPieces()
	return
}

func (mgr *ViewManager) DrawDynamic(screen *ebiten.Image) {
	if mgr.Piece2Put != nil {
		mgr.Piece2Put.Print(screen)
	}
	if mgr.PieceSel != nil {
		mgr.PieceSel.Print(screen)
	}
}

func (mgr *ViewManager) SetSelect(x, y int) {
	row, col := mgr.LocateCoord(x, y)
	mgr.SelIndex = mgr.GetIndex(row, col)
	mgr.TakePiece(row, col)
	mgr.PieceSel = mgr.Canvas.Piece[0].Take()
	mgr.Piece2Put = mgr.Canvas.Piece[SEL_SIDE].Take()
	mgr.Piece2Put.Put(x, y)
}

func (mgr *ViewManager) UpdateSelect(x, y int) {
	if mgr.Piece2Put != nil {
		mgr.Piece2Put.Put(mgr.GetCoord(mgr.LocateCoord(x, y)))
	}
	if mgr.PieceSel != nil {
		mgr.PieceSel.Put(x, y)
	}
}

func (mgr *ViewManager) ReleaseSelect(x, y int) {
	row, col := mgr.LocateCoord(x, y)
	mgr.PutPiece(row, col, 0)
	mgr.PieceSel = nil
	mgr.Piece2Put = nil
}
