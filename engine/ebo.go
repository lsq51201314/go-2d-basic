package engine

import (
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type EBO struct {
	id     uint32
	length int32
}

// 缓冲
func NewEbo(g *Object) EBO {
	var id uint32
	gl.GenBuffers(1, &id)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, id)
	return EBO{
		id: id,
	}
}

// 上传
func (e *EBO) Upload(arr []int32) {
	e.length = int32(len(arr))
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(arr)*int(INT_LENGTH), unsafe.Pointer(&arr[0]), gl.STATIC_DRAW)
}

// 绑定
func (e *EBO) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, e.id)
}

// 绘制
func (e *EBO) Draw() {
	gl.DrawElementsWithOffset(gl.TRIANGLES, e.length, gl.UNSIGNED_INT, 0)
}

// 销毁
func (e *EBO) Destroy() {
	gl.DeleteBuffers(1, &e.id)
}
