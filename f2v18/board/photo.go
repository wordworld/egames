package board

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Photo struct {
	*ebiten.Image
	*ebiten.DrawImageOptions
}

// 新建一张空白照片
func NewPhoto(width, heigt int) *Photo {
	return &Photo{
		Image:            ebiten.NewImage(width, heigt),
		DrawImageOptions: &ebiten.DrawImageOptions{},
	}
}

// 使用同一个 Image，位置不同
func TakePhoto(scene *ebiten.Image) *Photo {
	return &Photo{
		Image:            scene,
		DrawImageOptions: &ebiten.DrawImageOptions{},
	}
}
func (p *Photo) Take() *Photo {
	return TakePhoto(p.Image)
}

// 打印到 Image 上
func (p *Photo) Print(screen *ebiten.Image) {
	screen.DrawImage(p.Image, p.DrawImageOptions)
}

// 放到某个位置
func (p *Photo) Put(x, y int) *Photo {
	p.GeoM.Reset()
	return p.Mov(x, y)
}

// 挪动一个向量 [dx, dy]
func (p *Photo) Mov(dx, dy int) *Photo {
	p.GeoM.Translate(float64(dx), float64(dy))
	return p
}
