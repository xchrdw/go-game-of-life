package game

import (
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/mathgl/mgl32"
)

type Controls struct {
	Up       bool
	Down     bool
	Right    bool
	Left     bool
	In       bool
	Out      bool
	Scroll   float32
	Clicked  bool
	MousePos mgl32.Vec2
}

func (controls *Controls) KeyCallback(key glfw.Key, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Repeat {
		return
	}

	switch key {
	case glfw.KeyUp, glfw.KeyW:
		controls.Up = action == glfw.Press
	case glfw.KeyDown, glfw.KeyS:
		controls.Down = action == glfw.Press
	case glfw.KeyRight, glfw.KeyD:
		controls.Right = action == glfw.Press
	case glfw.KeyLeft, glfw.KeyA:
		controls.Left = action == glfw.Press
	case glfw.KeyX:
		controls.In = action == glfw.Press
	case glfw.KeyZ:
		controls.Out = action == glfw.Press
	}

}

func (controls *Controls) MouseButtonCallback(button glfw.MouseButton, action glfw.Action) {
	if button == glfw.MouseButton1 {
		controls.Clicked = action != glfw.Release
	}
}

func (controls *Controls) CursorCallback(position mgl32.Vec2) {
	controls.MousePos = position
}

func (controls *Controls) ScrollCallback(yoff float32) {
	controls.Scroll += yoff
}
