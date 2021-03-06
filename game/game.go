package game

import (
	"fmt"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/xchrdw/go-game-of-life/logic"
	"github.com/xchrdw/go-game-of-life/rendering"
	"github.com/xchrdw/mathext/f32math"
)

type Game struct {
	board       *logic.Board
	geometry    *rendering.BoardGeometry
	shader      *rendering.BoardShader
	texture     *rendering.BoardTexture
	boardUpdate float32
	interval    float32
	controls    *Controls
	camera      *Camera
}

func NewGame() *Game {
	board := logic.FromString(`
		XXXXXXXX
		X_XXXX_X
		XXXXXXXX
		`, 8, 3, 5)
	fmt.Printf("Board size: %vx%v, texture-size: %v\n", board.Width(), board.Height(), board.TextureSize())

	camera := NewCamera()
	controls := &Controls{}

	geometry := rendering.NewBoardGeometry()

	shader := rendering.NewBoardShader()

	texture := rendering.NewBoardTexture()

	texture.Update(board)
	texture.Activate(gl.TEXTURE0)

	return &Game{board, geometry, shader, texture, 0.0, 0.5, controls, camera}
}

func (g *Game) Update(deltaSec float32) {
	g.boardUpdate += deltaSec
	g.camera.updateCamera(deltaSec, g.controls)

	if g.boardUpdate > g.interval {
		g.board.NextGen()
		g.texture.Update(g.board)
		g.boardUpdate -= g.interval
		g.boardUpdate = f32math.Min(g.boardUpdate, g.interval)
	}
}

func (g *Game) Draw() {
	newView := g.camera.CalcViewMatrix()
	g.shader.SetView(newView)

	g.geometry.Bind()
	g.shader.Use()
	g.texture.Activate(gl.TEXTURE0)

	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
}

func (g *Game) ViewportChangedCallback(width int, height int) {
	g.camera.ViewportChangedCallback(width, height)
}

func (g *Game) KeyCallback(key glfw.Key, action glfw.Action, mods glfw.ModifierKey) {
	g.controls.KeyCallback(key, action, mods)
}
func (g *Game) MouseButtonCallback(button glfw.MouseButton, action glfw.Action) {
	g.controls.MouseButtonCallback(button, action)
}
func (g *Game) CursorPositionCallback(xpos float32, ypos float32) {
	g.controls.CursorCallback(mgl32.Vec2{xpos, ypos})
}

func (g *Game) ScrollCallback(yoff float32) {
	g.controls.ScrollCallback(yoff)
}

func (g *Game) Delete() {
	g.geometry.Delete()
	g.shader.Delete()
	g.texture.Delete()
}
