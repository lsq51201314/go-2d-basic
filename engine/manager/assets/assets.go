package assets

import (
	"os"
	"strings"

	"github.com/lsq51201314/go-2d/engine"
	gpk "github.com/lsq51201314/go-pack"
)

const resource = "./assets"                       //资源目录
const passwd = "yOLyovZoaB!wM0ieT8#j@d@!ZVN!6jY5" //包的密码

type Object struct {
	debug bool
	list  map[string]gpk.Object
}

// 资源
func NewAssets(g *engine.Object) Object {
	return Object{
		debug: false,
		list:  make(map[string]gpk.Object),
	}
}

// 调试
func (a *Object) SetDebug(debug bool) {
	a.debug = debug
}

// 打包
func (a *Object) Pack() bool {
	if dirs, err := getDirs(resource); err != nil {
		engine.Log("Assets.GetDirs", err)
		return false
	} else {
		for _, dir := range dirs {
			var g gpk.Object
			if err := g.CreateFromFolder(dir, dir+".gpk", passwd, process); err != nil {
				engine.Log("Assets.Gpk.CreateFromFolder", err)
				return false
			}
			g.Close()
		}
	}
	return true
}

// 载入
func (a *Object) Load() bool {
	if files, err := getFiles(resource); err != nil {
		engine.Log("Assets.GetFiles", err)
		return false
	} else {
		a.Destroy()
		for _, file := range files {
			var g gpk.Object
			if err := g.Load(file, passwd); err != nil {
				engine.Log("Assets.Gpk.Load", err)
				return false
			}
			name := strings.Replace(file, resource+"/", "", 1)
			name = strings.Replace(name, ".gpk", "", 1)
			a.list[name] = g
		}
	}
	return true
}

// 读取
func (a *Object) GetData(gpk string, name string) []byte {
	if a.debug {
		if data, err := os.ReadFile(resource + "/" + gpk + "/" + name); err != nil {
			engine.Log("Assets.os.ReadFile", err)
			return nil
		} else {
			return data
		}
	} else {
		g := a.list[gpk]
		if data, err := g.GetData(name); err != nil {
			engine.Log("Assets.Gpk.GetData", err)
			return nil
		} else {
			return data
		}
	}
}

// 销毁
func (a *Object) Destroy() {
	for _, v := range a.list {
		v.Close()
	}
	a.list = make(map[string]gpk.Object)
}
