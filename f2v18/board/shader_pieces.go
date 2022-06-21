//go:build ignore
// +build ignore

package main

// Uniform variables.
var Center vec2
var Radius float
var Antialias float
var Color vec4

// Fragment is the entry point of the fragment shader.
// Fragment returns the color value for the current position.
func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	pos := vec2(position.x, position.y)
	dist := distance(pos, Center) + Antialias/2
	t := smoothstep(0, Antialias/2, dist-Radius)
	return vec4(Color.rgb*(1-t), 1-t)
}
