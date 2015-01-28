package main

import (
	"fmt"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glh"
	"github.com/xchrdw/go-game-of-life/game"
	"math"

	"log"
	"os/exec"
	"runtime"
	"time"
)

func reexec() {
	log.Println("rerun")
	c := exec.Command("cmd", "/k", "start go run main.go")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	log.Fatal("restarted")
	return
}

func main() {
	// lock glfw/gl calls to a single thread
	runtime.LockOSThread()
	glfw.Init()
	defer glfw.Terminate()

	glfw.SetErrorCallback(func(code glfw.ErrorCode, desc string) {
		fmt.Println("GLFW ERROR! ", code)
		panic(desc)
	})

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 1) // 3
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	//glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)
	glfw.WindowHint(glfw.OpenglDebugContext, glfw.True)
	window, err := glfw.CreateWindow(800, 600, "Go Way of Life", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	gl.Init()

	gl.GetError() // THROW ERROR AWAY

	/*gl.DebugMessageCallback(func(source gl.GLenum, typ gl.GLenum, id uint, severity gl.GLenum, message string) {
		fmt.Println("===============")
		fmt.Println(message)
		fmt.Println("===============")
	})*/

	game := game.NewGame()
	defer game.Delete()

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		game.KeyCallback(w, key, action, mods)
	})
	window.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		gl.Viewport(0, 0, width, height)
	})

	last := glfw.GetTime()

	gl.ClearColor(0.3, 0.3, 0.3, 1.0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		now := glfw.GetTime()
		passed := now - last
		last = now

		game.Update(float32(passed))

		game.Draw()

		glh.OpenGLSentinel()
		window.SwapBuffers()
		gl.GetError() // THROW ERROR AWAY

		glfw.PollEvents()
		if window.GetKey(glfw.KeyR) == glfw.Press {
			reexec()
		}
		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			window.SetShouldClose(true)
		}
		glh.OpenGLSentinel()

		sleeptime := math.Max(0.0, 1.0/60.0-(glfw.GetTime()-now)) * float64(time.Second)
		time.Sleep(time.Duration(sleeptime))
	}
}
