package engine

import "github.com/go-gl/glfw/v3.3/glfw"

//销毁
func (g *Object) Destroy() {
	g.shader.Destroy()
	g.vao.Destroy()
	g.vbo.Destroy()
	g.ebo.Destroy()
	g.window.Destroy()
	glfw.Terminate()
}
