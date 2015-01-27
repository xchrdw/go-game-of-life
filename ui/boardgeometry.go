package ui

import (
	"encoding/binary"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"github.com/go-gl/mathgl/mgl32"
)

type Vertex struct {
	pos mgl32.Vec3
	uv  mgl32.Vec2
}

func v(x float32, y float32, z float32, u float32, v float32) Vertex {
	return Vertex{mgl32.Vec3{x, y, z}, mgl32.Vec2{u, v}}
}

const p float32 = 0.8

var verticies = []Vertex{
	v(-p, p, 0, 0, 0),
	v(-p, -p, 0, 0, 1),
	v(p, p, 0, 1, 0),
	v(p, -p, 0, 1, 1),
}

type BoardGeometry struct {
	vao gl.VertexArray
	vbo gl.Buffer
}

func NewBoardGeometry() BoardGeometry {
	defer glh.OpenGLSentinel()()

	vao := gl.GenVertexArray()
	vao.Bind()

	vbo := gl.GenBuffer()
	vbo.Bind(gl.ARRAY_BUFFER)

	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(verticies), verticies, gl.STATIC_DRAW)

	return BoardGeometry{vao, vbo}

}

func (g *BoardGeometry) Bind() {
	g.vao.Bind()
}

func (g *BoardGeometry) Delete() {
	g.vbo.Delete()
	g.vao.Delete()
}
