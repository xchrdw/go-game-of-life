package logic

import (
	"strings"

	"github.com/xchrdw/mathext/imath"
)

type Board struct {
	width       int
	height      int
	textureSize int
	cells       []bool
	lastGen     []bool
	texture     []byte
}

func CreateBoard(width int, height int) *Board {
	size := width * height
	texture_size := imath.NextPow2(imath.Max(width, height))
	b := Board{width, height, texture_size, make([]bool, size), make([]bool, size), make([]byte, texture_size*texture_size)}
	return &b
}

func FromString(board string, width int, height int, offset int) *Board {
	b := CreateBoard(width+offset*2, height+offset*2)
	if strings.Count(board, "X")+strings.Count(board, "_") != width*height {
		panic("invalid board size")
	}

	i := 0
	for _, c := range board {
		switch c {
		case 'X':
			pos := (i/width+offset)*b.width + (i%width + offset)
			b.cells[pos] = true
			i++
		case '_':
			i++
		}
	}
	if i != width*height {
		panic("string format is wrong")
	}
	b.updateTexture()
	return b
}

func (board *Board) Texture() []byte {
	return board.texture
}

func (board *Board) TextureSize() int {
	return board.textureSize
}

func (board *Board) Width() int {
	return board.width
}
func (board *Board) Height() int {
	return board.height
}

func (board *Board) Get(x int, y int) bool {
	if x < 0 || y < 0 || x >= board.width || y >= board.height {
		return false
	}
	return board.cells[x+y*board.width]
}

func (board *Board) Set(x int, y int, value bool) bool {
	old := board.cells[x+y*board.width]
	board.cells[x+y*board.width] = value
	return old
}

func (board *Board) NextGen() {
	for y := 0; y < board.height; y++ {
		for x := 0; x < board.width; x++ {
			living := board.livingNeighbours(x, y)
			alive := board.Get(x, y)
			if alive && (living == 2 || living == 3) {
				board.lastGen[x+y*board.width] = true
			} else if !alive && living == 3 {
				board.lastGen[x+y*board.width] = true
			} else {
				board.lastGen[x+y*board.width] = false

			}
		}
	}
	next := board.lastGen
	board.lastGen = board.cells
	board.cells = next

	board.updateTexture()
}

func (board *Board) updateTexture() {
	offsetx := (board.textureSize - board.width) / 2
	offsety := (board.textureSize - board.height) / 2
	for i, b := range board.cells {
		pos := (i/board.width+offsety)*board.textureSize + i%board.width + offsetx
		if b {
			board.texture[pos] = 255
		} else {
			board.texture[pos] = 0
		}
	}
}

func (board *Board) livingNeighbours(x int, y int) int {
	living := 0
	if board.Get(x-1, y-1) {
		living++
	}
	if board.Get(x, y-1) {
		living++
	}
	if board.Get(x+1, y-1) {
		living++
	}
	if board.Get(x-1, y) {
		living++
	}
	if board.Get(x+1, y) {
		living++
	}
	if board.Get(x-1, y+1) {
		living++
	}
	if board.Get(x, y+1) {
		living++
	}
	if board.Get(x+1, y+1) {
		living++
	}
	return living
}

func (board *Board) ToString() string {
	s := make([]byte, board.width*board.height+board.height-1)
	for i, b := range board.cells {
		var c byte
		if b {
			c = 'X'
		} else {
			c = '_'
		}
		s[i+i/board.width] = c

	}
	for i := 1; i < board.height; i++ {
		s[i*board.width+i-1] = ' '
	}
	return string(s)

}
