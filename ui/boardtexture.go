package ui

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"github.com/xchrdw/go-game-of-life/game"
)

type BoardTexture struct {
	texture gl.Texture
}

func NewBoardTexture() BoardTexture {
	defer glh.OpenGLSentinel()()
	texture := gl.GenTexture()
	texture.Bind(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	return BoardTexture{texture}
}

func (t *BoardTexture) Update(board *game.Board) {
	defer glh.OpenGLSentinel()()
	t.texture.Bind(gl.TEXTURE_2D)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RED, board.Width(), board.Height(), 0, gl.RED, gl.UNSIGNED_BYTE, board.Texture())
}

func (t *BoardTexture) Activate(slot gl.GLenum) {
	defer glh.OpenGLSentinel()()
	gl.ActiveTexture(slot)
	t.texture.Bind(gl.TEXTURE_2D)
}

func (t *BoardTexture) Delete() {
	t.texture.Delete()
}
