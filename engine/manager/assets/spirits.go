package assets

import (
	"sync"

	"github.com/lsq51201314/go-2d/engine"
)

type Spirits struct {
	game     *engine.Object
	resource *Object
	list     sync.Map
}

func NewSpirits(game *engine.Object, resource *Object) Spirits {
	return Spirits{
		game:     game,
		resource: resource,
		list:     sync.Map{},
	}
}

func (s *Spirits) Load(pack, name string) *engine.Texture {
	if tex, ok := s.list.Load(pack + "::" + name); ok {
		return tex.(*engine.Texture)
	}

	data := s.resource.GetData(pack, name)
	if len(data) == 0 {
		return nil
	}
	if tex, ok := engine.NewTexture(s.game, data); !ok {
		return nil
	} else {
		s.list.Store(pack+"::"+name, &tex)
		return &tex
	}
}

// 销毁
func (s *Spirits) Destroy() {
	s.list.Range(func(key, value any) bool {
		texture := value.(*engine.Texture)
		texture.Destroy()
		return true
	})
	s.list = sync.Map{}
}
