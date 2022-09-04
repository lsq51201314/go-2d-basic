package engine

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type CloseCallback func(g *Object)
type SizeCallback func(g *Object, width int, height int)
type KeyCallback func(g *Object, key Key, scancode int, action Action, mods ModifierKey)
type MouseButtonCallback func(g *Object, button MouseButton, action Action, mods ModifierKey)
type CursorPosCallback func(g *Object, xpos float64, ypos float64)
type ScrollCallback func(g *Object, xoff float64, yoff float64)

//关闭回调
func (g *Object) SetCloseCallback(cfun CloseCallback) {
	g.closeCb = cfun
}

//关闭回调
func (g *Object) closeCallback(w *glfw.Window) {
	if g.closeCb != nil {
		g.closeCb(g)
	}
}

//调整大小回调
func (g *Object) SetSizeCallback(cfun SizeCallback) {
	g.sizeCb = cfun
}

//调整大小回调
func (g *Object) sizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	if g.sizeCb != nil {
		g.sizeCb(g, width, height)
	}
}

//按键回调
func (g *Object) SetKeyCallback(cfun KeyCallback) {
	g.keyCb = cfun
}

//按键回调
func (g *Object) keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if g.keyCb != nil {
		g.keyCb(g, Key(key), scancode, Action(action), ModifierKey(mods))
	}
	// fmt.Println(key, scancode, action, mods)
	switch key {
	case glfw.KeyEscape:
		if action == glfw.Press {
			g.closeCallback(w)
		}
	}
}

//鼠标回调
func (g *Object) SetMouseButtonCallback(cfun MouseButtonCallback) {
	g.mouseButtonCb = cfun
}

//鼠标回调
func (g *Object) mouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if g.mouseButtonCb != nil {
		g.mouseButtonCb(g, MouseButton(button), Action(action), ModifierKey(mods))
	}
}

//鼠标指针
func (g *Object) SetCursorPosCallback(cfun CursorPosCallback) {
	g.cursorPosCb = cfun
}

//鼠标指针
func (g *Object) cursorPosCallback(w *glfw.Window, xpos float64, ypos float64) {
	if g.cursorPosCb != nil {
		g.cursorPosCb(g, xpos, ypos)
	}
}

//鼠标滚动
func (g *Object) SetScrollCallback(cfun ScrollCallback) {
	g.scrollCb = cfun
}

//鼠标滚动
func (g *Object) scrollCallback(w *glfw.Window, xoff float64, yoff float64) {
	if g.scrollCb != nil {
		g.scrollCb(g, xoff, yoff)
	}
}
