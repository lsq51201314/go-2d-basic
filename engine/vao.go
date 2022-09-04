package engine

import "github.com/go-gl/gl/v3.3-core/gl"

type VAO struct {
	id uint32
}

// 顶点
func NewVao(g *Object) VAO {
	var id uint32
	gl.GenVertexArrays(1, &id)
	gl.BindVertexArray(id)
	return VAO{
		id: id,
	}
}

// 绑定
func (v *VAO) Bind() {
	gl.BindVertexArray(v.id)
}

// 销毁
func (v *VAO) Destroy() {
	gl.DeleteVertexArrays(1, &v.id)
}
