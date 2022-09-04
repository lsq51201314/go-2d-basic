package engine

type Camera struct {
	game          *Object
	focusPosition Vec2
	zoom          float32
}

func NewCamera(game *Object) Camera {
	w, h := game.window.GetSize()
	return Camera{
		game: game,
		focusPosition: Vec2{
			X: float32(w) / 2,
			Y: float32(h) / 2,
		},
		zoom: 1,
	}
}

func (c *Camera) SetZoom(zoom float32) {
	c.zoom = zoom
}

func (c *Camera) GetZoom() float32 {
	return c.zoom
}

func (c *Camera) SetPosition(x, y float32) {
	c.focusPosition.X = x
	c.focusPosition.Y = y
}

func (c *Camera) GetPosition() Vec2 {
	return c.focusPosition
}

func (c *Camera) GetProjectionMatrix() Mat4 {
	w, h := c.game.window.GetSize()
	left := c.focusPosition.X - float32(w)/2
	right := c.focusPosition.X + float32(w)/2
	top := c.focusPosition.Y - float32(h)/2
	bottom := c.focusPosition.Y + float32(h)/2
	orthoMatrix := CreateOrthographicOffCenter(left, right, bottom, top, 0.01, 100.0)
	zoomMatrix := CreateScale1f(c.zoom)
	return Multiply(orthoMatrix, zoomMatrix)
}
