package board

// 选项列表
type Option struct {
	*DisposableOption
	// size
	Width  int
	Height int
	// scale
	ScaleX float64
	ScaleY float64
	// pos
	X int
	Y int
	// grids
	Rows int
	Cols int
}

type Operator func(opt *Option)

func (op *Option) Apply(options ...Operator) {
	op.Disposable()
	for _, opr := range options {
		opr(op)
	}
}

func WithSize(width, height int) Operator {
	return func(opt *Option) {
		opt.Width = width
		opt.Height = height
	}
}

func WidthScale(x, y float64) Operator {
	return func(opt *Option) {
		opt.ScaleX = x
		opt.ScaleY = y
	}
}

func WidthPos(x, y int) Operator {
	return func(opt *Option) {
		opt.X = x
		opt.Y = y
	}
}

func WithGrid(rols, cols int) Operator {
	return func(opt *Option) {
		opt.Rows = rols
		opt.Cols = cols
	}
}

// 一次性的选项
type DisposableOption struct {
	UpdateAll bool // 全量重绘
}

func (op *Option) Disposable() { // 重置一次性选项
	op.DisposableOption = new(DisposableOption)
}

func WidthUpdateAll() Operator {
	return func(opt *Option) {
		opt.UpdateAll = true
	}
}
