package game

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/xchrdw/mathext/f32math"
)

type Camera struct {
	pos  mgl32.Vec2
	zoom float32
}

func (camera *Camera) CalcViewMatrix() mgl32.Mat4 {
	return mgl32.Translate3D(camera.pos.X(), camera.pos.Y(), 0).Mul4(mgl32.Scale3D(camera.zoom, camera.zoom*4/3, 1.0))
}

func (camera *Camera) updateCamera(deltaSec float32, controls Controls) {
	movement := mgl32.Vec2{}
	if controls.Up {
		movement[1] += +1
	}
	if controls.Down {
		movement[1] += -1
	}
	if controls.Right {
		movement[0] += +1
	}
	if controls.Left {
		movement[0] += -1
	}
	if movement.Len() > 0 {
		movement = movement.Normalize()
	}
	movement = movement.Mul(deltaSec)
	camera.pos = camera.pos.Add(movement)

	var zoom float32
	if controls.In {
		zoom += 1
	}
	if controls.Out {
		zoom -= 1
	}
	camera.zoom += zoom * deltaSec
	camera.zoom = f32math.Max(0.1, camera.zoom)
}
