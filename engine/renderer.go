package engine

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// 渲染开始
func (g *Object) RendererBegan() {
	g.deltaTime = glfw.GetTime() - g.totalElapsedSeconds
	g.totalElapsedSeconds = glfw.GetTime()
	//fmt.Println(1 / g.deltaTime) //FPS
	//处理事件
	glfw.PollEvents()
	//清屏
	gl.ClearColor(0.0, 0.0, 0.0, 1.0) //填充颜色
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	//摄像机
	g.shader.UploadMat4("uProjection", g.camera.GetProjectionMatrix())
}

// 渲染结束
func (g *Object) RendererEnd() {
	//交换缓存
	g.window.SwapBuffers()
}
