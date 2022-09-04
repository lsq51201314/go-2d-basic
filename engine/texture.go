package engine

import (
	"bytes"
	"errors"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Texture struct {
	game   *Object
	id     uint32
	width  int32
	height int32
}

// 新建纹理
func NewTexture(game *Object, data []byte) (texture Texture, ok bool) {
	reader := bytes.NewReader(data)
	img, _, err := image.Decode(reader)
	if err != nil {
		Log("NewTexture.Image.Decode", err)
		return
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		Log("NewTexture.Image.NewRGBA", errors.New("length mismatch"))
		return
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	texture = Texture{
		game:   game,
		width:  int32(rgba.Rect.Size().X),
		height: int32(rgba.Rect.Size().Y),
	}

	var id uint32
	gl.GenTextures(1, &id)
	gl.BindTexture(gl.TEXTURE_2D, id)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		texture.width,
		texture.height,
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	texture.id = id
	ok = true
	return
}

// 绘制
func (t *Texture) Draw(x, y, scale, angle float32, horizontal, vertical bool) {
	trans := CreateTranslation(x, y, 0)
	sca := CreateScale3f(float32(t.width)*scale, float32(t.height)*scale, 1)
	rot := CreateRotationZ(float32(angle))

	mul := Multiply(sca, rot)
	mul = Multiply(mul, trans)
	if vertical {
		mul = RotationX1f(mul, 180)
	}
	if horizontal {
		mul = RotationY1f(mul, 180)
	}

	t.game.shader.UploadMat4("uView", mul)

	t.game.shader.UploadTexture("TEX_SAMPLER", 0)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, t.id)

	t.game.ebo.Draw()
}

// 销毁
func (t *Texture) Destroy() {
	gl.DeleteTextures(1, &t.id)
}

// 获取图像
func (t *Texture) GetPixel(x, y float64) Color {
	if x >= 0 && int32(x) <= t.width && y >= 0 && int32(y) <= t.height {
		gl.BindTexture(gl.TEXTURE_2D, t.id)
		pixels := make([]byte, 4)
		gl.GetTextureSubImage(
			t.id,
			0,
			int32(x),
			int32(y),
			0,
			1,
			1,
			1,
			gl.RGBA,
			gl.UNSIGNED_BYTE,
			int32(len(pixels)),
			unsafe.Pointer(&pixels[0]),
		)
		return Color{
			R: pixels[0],
			G: pixels[1],
			B: pixels[2],
			A: pixels[3],
		}
	}



	return Color{}
}
