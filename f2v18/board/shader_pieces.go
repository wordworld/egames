//go:build ignore
// +build ignore

package main

// Uniform variables.
var Center vec2
var Radius float
var Antialias float
var Color vec4
var ShadowColor vec4
var ShadowDist float

// Fragment is the entry point of the fragment shader.
// Fragment returns the color value for the current position.
func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	pos := vec2(position.x, position.y)
	dist := distance(pos, Center) + Antialias/2
	t := smoothstep(0, Antialias/2, dist-Radius)
	delta := Color - ShadowColor
	view := ShadowColor + delta*smoothstep(Radius-ShadowDist, Radius*0.95, Radius-dist)
	return view * (1 - t) // 消除锯齿
}
