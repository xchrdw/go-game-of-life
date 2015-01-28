package game

import (
	glfw "github.com/go-gl/glfw3"
)

type Controls struct {
	Up    bool
	Down  bool
	Right bool
	Left  bool
	In    bool
	Out   bool
}

func (controls *Controls) KeyCallback(w *glfw.Window, key glfw.Key, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Repeat {
		return
	}

	switch key {
	case glfw.KeyUp:
		controls.Up = action == glfw.Press
	case glfw.KeyDown:
		controls.Down = action == glfw.Press
	case glfw.KeyRight:
		controls.Right = action == glfw.Press
	case glfw.KeyLeft:
		controls.Left = action == glfw.Press
	case glfw.KeyX:
		controls.In = action == glfw.Press
	case glfw.KeyZ:
		controls.Out = action == glfw.Press
	}
}
