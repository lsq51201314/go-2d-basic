package engine

type Animation struct {
	game      *Object
	width     float32
	height    float32
	center    Vec2
	frames    []*Texture
	index     int
	deltaTime float64
}

func NewAnimation(game *Object, W, H, centerX, centerY float32) Animation {
	return Animation{
		game:   game,
		width:  W,
		height: H,
		center: Vec2{X: centerX, Y: centerY},
		frames: make([]*Texture, 0),
		index:  -1,
	}
}

func (a *Animation) AddFrame(texture *Texture) bool {
	a.frames = append(a.frames, texture)
	return true
}

func (a *Animation) Renderer(x, y, scale, angle float32, start, end, fps int, horizontal, vertical bool) {
	if start < 0 || start > len(a.frames)-1 {
		start = 0
	}
	if end < 0 || end > len(a.frames)-1 {
		end = 0
	}
	if end < start {
		end = start
	}
	if a.index == -1 {
		a.index = start
	}

	tmp := 1 / float64(fps)
	a.deltaTime += a.game.deltaTime
	if a.deltaTime >= tmp {
		a.index++
		a.deltaTime = 0
	}

	if a.index >= end {
		a.index = start
	}
	a.frames[a.index].Draw(x, y, scale, angle, horizontal, vertical)
}
