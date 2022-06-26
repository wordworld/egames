package board

import "f2v18/template/gen/tint"

// 所在棋盘位置(行、列)相关逻辑的接口
type Location interface {
	LocateCoord(x, y int) (int, int)  // 用坐标 (x,y) 定位(row,col)
	LocateIndex(index int) (int, int) // 用索引 index 定位(row,col)
	GetCoord(row, col int) (int, int) // 计算位置(row,col)的坐标 (x,y)
	GetIndex(row, col int) int        // 计算位置(row,col)的索引 index
	Quad(row, col int) *tint.Quad
	GetLeft() int
	GetTop() int
	GetWidth() int
	GetHeight() int
	GetRight() int
	GetBottom() int
}
