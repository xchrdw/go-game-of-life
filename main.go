package main

import (
	"fmt"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glh"
	"github.com/xchrdw/go-game-of-life/ui"

	"runtime"
	"time"
)

func main() {
	// lock glfw/gl calls to a single thread
	runtime.LockOSThread()
	glfw.Init()
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)
	window, err := glfw.CreateWindow(800, 600, "Go Way of Life", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	gl.Init()

	gl.GetError() // THROW ERROR AWAY

	gl.DebugMessageCallback(func(source gl.GLenum, typ gl.GLenum, id uint, severity gl.GLenum, message string) {
		fmt.Println("===============")
		fmt.Println(message)
		fmt.Println("===============")
	})

	game := ui.NewGame()
	defer game.Delete()

	last := time.Now()

	gl.ClearColor(0.3, 0.3, 0.3, 1.0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		now := time.Now()
		passed := now.Sub(last).Seconds()
		last = now

		game.Update(float32(passed))

		game.Draw()

		glh.OpenGLSentinel()
		window.SwapBuffers()
		gl.GetError() // THROW ERROR AWAY

		glfw.PollEvents()
		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			window.SetShouldClose(true)
		}

		glh.OpenGLSentinel()
		time.Sleep(10 * time.Millisecond)
	}
}
