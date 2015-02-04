package game

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/xchrdw/mathext/f32math"
)

type Camera struct {
	pos       mgl32.Vec2
	zoom      float32
	viewPort  mgl32.Vec2
	lastMouse mgl32.Vec2
}

func NewCamera() *Camera {
	return &Camera{mgl32.Vec2{}, 1.0, mgl32.Vec2{}, mgl32.Vec2{}}
}

func (camera *Camera) CalcViewMatrix() mgl32.Mat4 {
	aspectRatio := camera.viewPort.X() / camera.viewPort.Y()

	return mgl32.Translate3D(camera.pos.X(), camera.pos.Y(), 0).Mul4(
		mgl32.HomogRotate3D(0.0, mgl32.Vec3{0.0, 0.0, 1.0})).Mul4(
		mgl32.Scale3D(camera.zoom, camera.zoom*aspectRatio, 1.0))
}

func (camera *Camera) ViewportChangedCallback(w int, h int) {
	camera.viewPort = mgl32.Vec2{float32(w), float32(h)}
}

func (camera *Camera) updateCamera(deltaSec float32, controls *Controls) {
	movement := mgl32.Vec2{}
	if controls.Up {
		movement[1] += 1
	}
	if controls.Down {
		movement[1] -= 1
	}
	if controls.Right {
		movement[0] += 1
	}
	if controls.Left {
		movement[0] -= 1
	}
	if movement.Len() > 0 {
		movement = movement.Normalize()
	}
	camera.pos = camera.pos.Add(movement.Mul(deltaSec))

	mouseDiff := controls.MousePos.Sub(camera.lastMouse).Mul(2)
	mouseDiff[0] /= camera.viewPort[0]
	mouseDiff[1] /= -camera.viewPort[1]
	if controls.Clicked && mouseDiff.Len() > 0 {
		camera.pos = camera.pos.Add(mouseDiff.Mul(1.0))
	}

	var zoom float32
	if controls.In {
		zoom += 1
	}
	if controls.Out {
		zoom -= 1
	}
	if controls.Scroll > 0.05 {
		zoom += controls.Scroll * 3
		controls.Scroll /= 1.7
	}
	if controls.Scroll < -0.05 {
		zoom += controls.Scroll * 3
		controls.Scroll /= 1.7
	}
	camera.zoom += zoom * deltaSec
	camera.zoom = f32math.Max(0.05, camera.zoom)

	camera.lastMouse = controls.MousePos
}
