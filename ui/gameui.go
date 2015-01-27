package ui

import (
	"github.com/go-gl/gl"
	"github.com/xchrdw/go-game-of-life/game"
)

type Game struct {
	board       *game.Board
	geometry    BoardGeometry
	shader      BoardShader
	texture     BoardTexture
	boardUpdate float32
	interval    float32
}

func NewGame() *Game {
	board := game.FromString("__X_ XX__ __XX _X__", 4, 4)

	geometry := NewBoardGeometry()

	shader := NewBoardShader()

	texture := NewBoardTexture()
	texture.Update(board)
	texture.Activate(gl.TEXTURE0)

	return &Game{board, geometry, shader, texture, 0.0, 0.5}
}

func (g *Game) Update(deltaSec float32) {
	g.boardUpdate += deltaSec
	if g.boardUpdate > g.interval {
		g.board.NextGen()
		g.texture.Update(g.board)
		g.boardUpdate -= g.interval
	}
}

func (g *Game) Draw() {
	g.geometry.Bind()
	g.shader.Use()
	g.texture.Activate(gl.TEXTURE0)

	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
}

func (g *Game) Delete() {
	g.geometry.Delete()
	g.shader.Delete()
	g.texture.Delete()
}
