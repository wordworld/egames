package view

const (
	LIGHT_SIDE = 0 // 白子
	DARK_SIDE  = 1 // 黑子
	SEL_SIDE   = 2 // 选中的棋子
)

func (cvs *ViewManager) DrawPieces() {
	// 2
	for _, col := range []int{(cvs.Cols - 1) / 2, cvs.Cols / 2} {
		cvs.PutPiece(cvs.Rows-1, col, DARK_SIDE)
	}
	// 18
	for row := 0; row < 3; row++ {
		for col := 0; col < cvs.Cols; col++ {
			cvs.PutPiece(row, col, LIGHT_SIDE)
		}
	}
}
