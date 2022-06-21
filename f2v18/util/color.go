package util

import "image/color"

func Color2Vec(rgba color.Color) []float32 {
	r, g, b, a := rgba.RGBA()
	return []float32{float32(r&0xff) / 255, float32(g&0xff) / 255, float32(b&0xff) / 255, float32(a&0xff) / 255}
}

func InvColor(rgba color.Color) color.Color {
	r, g, b, a := rgba.RGBA()
	return &color.RGBA{255 - uint8(r&0xff), 255 - uint8(g&0xff), 255 - uint8(b&0xff), uint8(a & 0xff)}
}
