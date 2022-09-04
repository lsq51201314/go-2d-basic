package engine

import (
	"errors"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Shader struct {
	id uint32
}

// 着色器
func NewShader(vertex, fragment string) (shader Shader, ok bool) {
	//顶点着色器
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	vertexString, vertexFree := gl.Strs(vertex + "\x00")
	defer vertexFree()
	gl.ShaderSource(vertexShader, 1, vertexString, nil)
	gl.CompileShader(vertexShader)
	var vertexSuccess int32
	vertexInfoLog := make([]uint8, 512)
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &vertexSuccess)
	if vertexSuccess == gl.FALSE {
		gl.GetShaderInfoLog(vertexShader, int32(len(vertexInfoLog)), nil, &vertexInfoLog[0])
		Log("NewShader.VertexShader", errors.New(string(vertexInfoLog)))
		return
	}
	//片段着色器
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fragmentString, fragmentFree := gl.Strs(fragment + "\x00")
	defer fragmentFree()
	gl.ShaderSource(fragmentShader, 1, fragmentString, nil)
	gl.CompileShader(fragmentShader)
	var fragmentSuccess int32
	fragmentInfoLog := make([]uint8, 512)
	gl.GetShaderiv(fragmentShader, gl.COMPILE_STATUS, &fragmentSuccess)
	if fragmentSuccess == gl.FALSE {
		gl.GetShaderInfoLog(vertexShader, int32(len(fragmentInfoLog)), nil, &fragmentInfoLog[0])
		Log("NewShader.FragmentShader", errors.New(string(fragmentInfoLog)))
		return
	}
	//着色器程序
	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)
	var shaderSuccess int32
	shaderInfoLog := make([]uint8, 512)
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &shaderSuccess)
	if shaderSuccess == gl.FALSE {
		gl.GetProgramInfoLog(shaderProgram, int32(len(shaderInfoLog)), nil, &shaderInfoLog[0])
		Log("NewShader.ShaderProgram", errors.New(string(shaderInfoLog)))
		return
	}
	//删除着色器
	gl.DetachShader(shaderProgram, vertexShader)
	gl.DetachShader(shaderProgram, fragmentShader)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	//返回
	shader = Shader{
		id: shaderProgram,
	}
	ok = true
	return
}

// 使用
func (s *Shader) Use() {
	gl.UseProgram(s.id)
}

// 上传
func (s *Shader) UploadMat4(name string, mat4 Mat4) {
	nameStr := gl.Str(name + "\x00")
	vertLocation := gl.GetUniformLocation(s.id, nameStr)
	gl.UseProgram(s.id)
	gl.UniformMatrix4fv(vertLocation, 1, false, &mat4[0][0])
}

// 上传
func (s *Shader) UploadTexture(name string, slot int32) {
	nameStr := gl.Str(name + "\x00")
	vertLocation := gl.GetUniformLocation(s.id, nameStr)
	gl.UseProgram(s.id)
	gl.Uniform1i(vertLocation, slot)
}

// 销毁
func (s *Shader) Destroy() {
	gl.DeleteProgram(s.id)
}
