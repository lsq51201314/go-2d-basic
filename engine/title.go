package engine

import (
	"bytes"
	"image"
	"image/png"
)

//设置标题
func (g *Object) SetTitle(str string) {
	g.window.SetTitle(str)
}

//设置图标
func (g *Object) SetIcon(data []byte) bool {
	buf := bytes.NewBuffer(data)
	if img, err := png.Decode(buf); err != nil {
		Log("png.Decode", err)
		return false
	} else {
		g.window.SetIcon([]image.Image{img})
		return true
	}
}
