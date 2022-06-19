//go:genrate -command file2byteslice ../tools/bin/file2byteslice
//go:generate file2byteslice -input board/shader_pieces.go -output board/shader_pieces_var.go -package board -var shader_pieces

package main

import (
	"f2v18/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
	g.Quit()
}
