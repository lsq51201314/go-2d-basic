package engine

import (
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type VBO struct {
	id uint32
}

// 缓存
func NewVbo(g *Object) VBO {
	var id uint32
	gl.GenBuffers(1, &id)
	gl.BindBuffer(gl.ARRAY_BUFFER, id)
	return VBO{
		id: id,
	}
}

// 上传
func (v *VBO) Upload(arr []float32) {
	gl.BufferData(gl.ARRAY_BUFFER, len(arr)*int(FLOAT_LENGTH), unsafe.Pointer(&arr[0]), gl.STATIC_DRAW)
}

// 绑定
func (v *VBO) Bind(location uint32, size, length, offset int32) {
	gl.VertexAttribPointerWithOffset(location, size, gl.FLOAT, false, length*FLOAT_LENGTH, uintptr(offset*FLOAT_LENGTH))
	gl.EnableVertexAttribArray(location)
}

// 销毁
func (v *VBO) Destroy() {
	gl.DeleteBuffers(1, &v.id)
}
