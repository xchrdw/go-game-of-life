package rendering

import (
	"encoding/binary"
	"unsafe"

	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"github.com/go-gl/mathgl/mgl32"
)

const vertex = `#version 330
in vec3 v_position;
in vec2 v_uv;

out vec2 f_uv;

uniform mat4 view;

void main() {
    gl_Position = view * vec4(v_position, 1.0);
    f_uv = v_uv ;
}`

const fragment = `#version 330
in vec2 f_uv;
out vec4 outColor;

uniform sampler2D board;
void main() {
    float r = texture(board, f_uv).r;
    outColor = vec4(f_uv.x*0.3, f_uv.y*0.3, r, 1.0);
}`

type BoardShader struct {
	program    gl.Program
	viewMatLoc gl.UniformLocation
}

func NewBoardShader() *BoardShader {
	defer glh.OpenGLSentinel()()

	vertex_shader := glh.Shader{gl.VERTEX_SHADER, vertex}
	fragment_shader := glh.Shader{gl.FRAGMENT_SHADER, fragment}

	program := glh.NewProgram(vertex_shader, fragment_shader)
	program.BindFragDataLocation(0, "outColor")
	program.Use()

	var v Vertex
	positionAttrib := program.GetAttribLocation("v_position")
	positionAttrib.AttribPointer(3, gl.FLOAT, false, binary.Size(v), unsafe.Offsetof(v.pos))
	positionAttrib.EnableArray()

	uvAttrib := program.GetAttribLocation("v_uv")
	uvAttrib.AttribPointer(2, gl.FLOAT, false, binary.Size(v), unsafe.Offsetof(v.uv))
	uvAttrib.EnableArray()

	textureLoc := program.GetUniformLocation("board")
	textureLoc.Uniform1i(0)

	viewMatLoc := program.GetUniformLocation("view")
	viewMatLoc.UniformMatrix4fv(false, mgl32.Ident4())

	return &BoardShader{program, viewMatLoc}
}

func (s *BoardShader) Use() {
	s.program.Use()
}

func (s *BoardShader) SetView(view mgl32.Mat4) {
	s.viewMatLoc.UniformMatrix4fv(false, view)
}

func (s *BoardShader) Delete() {
	s.program.Delete()
}
