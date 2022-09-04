package main

import (
	"fmt"
	"runtime"

	"github.com/lsq51201314/go-2d/engine"
	"github.com/lsq51201314/go-2d/engine/manager/assets"
)

func init() {
	runtime.LockOSThread()
}

var resource assets.Object
var spirits assets.Spirits
var tex engine.Texture
var game engine.Object

type Vec2 struct {
	X, Y float32
}

func main() {
	if game.Init(1920, 1080) {
		resource = assets.NewAssets(&game)
		spirits = assets.NewSpirits(&game, &resource)
		defer resource.Destroy()
		defer game.Destroy()
		defer spirits.Destroy()
		//resource.Pack()
		resource.Load()
		resource.SetDebug(true) //调试模式
		game.SetTitle("我的游戏")
		game.SetIcon(resource.GetData("config", "title.png"))
		game.SetCloseCallback(closeCallback)
		game.SetSizeCallback(sizeCallback)
		game.SetKeyCallback(keyCallback)
		game.SetMouseButtonCallback(mouseButtonCallback)
		game.SetCursorPosCallback(cursorPosCallback)
		game.SetScrollCallback(scrollCallback)

		tex, _ = engine.NewTexture(&game, resource.GetData("config", "uvgrid.jpg"))

		stand1 := engine.NewAnimation(&game, 277, 277, 140, 140)
		stand2 := engine.NewAnimation(&game, 277, 277, 140, 140)
		stand3 := engine.NewAnimation(&game, 277, 277, 140, 140)
		stand4 := engine.NewAnimation(&game, 277, 277, 140, 140)
		stand5 := engine.NewAnimation(&game, 277, 277, 140, 140)
		stand6 := engine.NewAnimation(&game, 277, 277, 140, 140)
		stand7 := engine.NewAnimation(&game, 277, 277, 140, 140)
		stand8 := engine.NewAnimation(&game, 277, 277, 140, 140)

		run1 := engine.NewAnimation(&game, 277, 277, 140, 140)
		run2 := engine.NewAnimation(&game, 277, 277, 140, 140)
		run3 := engine.NewAnimation(&game, 277, 277, 140, 140)
		run4 := engine.NewAnimation(&game, 277, 277, 140, 140)
		run5 := engine.NewAnimation(&game, 277, 277, 140, 140)
		run6 := engine.NewAnimation(&game, 277, 277, 140, 140)
		run7 := engine.NewAnimation(&game, 277, 277, 140, 140)
		run8 := engine.NewAnimation(&game, 277, 277, 140, 140)

		attack1 := engine.NewAnimation(&game, 277, 277, 140, 140)
		attack2 := engine.NewAnimation(&game, 277, 277, 140, 140)
		attack3 := engine.NewAnimation(&game, 277, 277, 140, 140)
		attack4 := engine.NewAnimation(&game, 277, 277, 140, 140)
		attack5 := engine.NewAnimation(&game, 277, 277, 140, 140)
		attack6 := engine.NewAnimation(&game, 277, 277, 140, 140)
		attack7 := engine.NewAnimation(&game, 277, 277, 140, 140)
		attack8 := engine.NewAnimation(&game, 277, 277, 140, 140)

		die1 := engine.NewAnimation(&game, 277, 277, 140, 140)
		die2 := engine.NewAnimation(&game, 277, 277, 140, 140)
		die3 := engine.NewAnimation(&game, 277, 277, 140, 140)
		die4 := engine.NewAnimation(&game, 277, 277, 140, 140)
		die5 := engine.NewAnimation(&game, 277, 277, 140, 140)
		die6 := engine.NewAnimation(&game, 277, 277, 140, 140)
		die7 := engine.NewAnimation(&game, 277, 277, 140, 140)
		die8 := engine.NewAnimation(&game, 277, 277, 140, 140)

		for i := 0; i < 145; i++ {
			stand1.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			stand2.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			stand3.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			stand4.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			stand5.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			stand6.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			stand7.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			stand8.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))

			run1.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			run2.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			run3.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			run4.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			run5.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			run6.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			run7.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			run8.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))

			attack1.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			attack2.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			attack3.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			attack4.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			attack5.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			attack6.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			attack7.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			attack8.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))

			die1.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			die2.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			die3.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			die4.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			die5.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			die6.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			die7.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
			die8.AddFrame(spirits.Load("role", fmt.Sprintf("%04d", i)+".png"))
		}

		for game.GetRunning() {
			game.RendererBegan()
			//*************************************************
			tex.Draw(512, 512, 1, 0, false, false)

			stand1.Renderer(200, 200, 1, 0, 0, 6, 8, false, false)
			stand2.Renderer(400, 200, 1, 0, 7, 13, 8, false, false)
			stand3.Renderer(600, 200, 1, 0, 14, 20, 8, false, false)
			stand4.Renderer(800, 200, 1, 0, 21, 27, 8, false, false)
			stand5.Renderer(1000, 200, 1, 0, 28, 34, 8, false, false)
			stand6.Renderer(1200, 200, 1, 0, 7, 13, 8, true, false)
			stand7.Renderer(1400, 200, 1, 0, 14, 20, 8, true, false)
			stand8.Renderer(1600, 200, 1, 0, 21, 27, 8, true, false)

			run1.Renderer(200, 400, 1, 0, 35, 43, 10, false, false)
			run2.Renderer(400, 400, 1, 0, 44, 52, 10, false, false)
			run3.Renderer(600, 400, 1, 0, 53, 61, 10, false, false)
			run4.Renderer(800, 400, 1, 0, 62, 70, 10, false, false)
			run5.Renderer(1000, 400, 1, 0, 71, 79, 10, false, false)
			run6.Renderer(1200, 400, 1, 0, 44, 52, 10, true, false)
			run7.Renderer(1400, 400, 1, 0, 53, 61, 10, true, false)
			run8.Renderer(1600, 400, 1, 0, 62, 70, 10, true, false)

			attack1.Renderer(200, 600, 1, 0, 80, 88, 10, false, false)
			attack2.Renderer(400, 600, 1, 0, 89, 97, 10, false, false)
			attack3.Renderer(600, 600, 1, 0, 98, 106, 10, false, false)
			attack4.Renderer(800, 600, 1, 0, 107, 115, 10, false, false)
			attack5.Renderer(1000, 600, 1, 0, 116, 124, 10, false, false)
			attack6.Renderer(1200, 600, 1, 0, 89, 97, 10, true, false)
			attack7.Renderer(1400, 600, 1, 0, 98, 106, 10, true, false)
			attack8.Renderer(1600, 600, 1, 0, 107, 115, 10, true, false)

			die1.Renderer(200, 800, 1, 0, 125, 128, 4, false, false)
			die2.Renderer(400, 800, 1, 0, 129, 132, 4, false, false)
			die3.Renderer(600, 800, 1, 0, 133, 136, 4, false, false)
			die4.Renderer(800, 800, 1, 0, 137, 140, 4, false, false)
			die5.Renderer(1000, 800, 1, 0, 141, 144, 4, false, false)
			die6.Renderer(1200, 800, 1, 0, 129, 132, 4, true, false)
			die7.Renderer(1400, 800, 1, 0, 133, 136, 4, true, false)
			die8.Renderer(1600, 800, 1, 0, 137, 140, 4, true, false)
			//*************************************************
			game.RendererEnd()
		}
	}
}

func closeCallback(g *engine.Object) {
	g.SetRunning(false)
}

func sizeCallback(g *engine.Object, width int, height int) {

}

func keyCallback(g *engine.Object, key engine.Key, scancode int, action engine.Action, mods engine.ModifierKey) {
	if key == engine.KeyEscape {
		if action == engine.Press {
			closeCallback(g)
		}
	}

	if key == engine.KeyW {
		c := game.GetCamera()
		p := c.GetPosition()
		p.Y -= 20
		c.SetPosition(p.X, p.Y)
	}

	if key == engine.KeyS {
		c := game.GetCamera()
		p := c.GetPosition()
		p.Y += 20
		c.SetPosition(p.X, p.Y)
	}

	if key == engine.KeyA {
		c := game.GetCamera()
		p := c.GetPosition()
		p.X -= 20
		c.SetPosition(p.X, p.Y)
	}

	if key == engine.KeyD {
		c := game.GetCamera()
		p := c.GetPosition()
		p.X += 20
		c.SetPosition(p.X, p.Y)
	}
}

func mouseButtonCallback(g *engine.Object, button engine.MouseButton, action engine.Action, mods engine.ModifierKey) {

}

func cursorPosCallback(g *engine.Object, xpos float64, ypos float64) {
	g.SetTitle(fmt.Sprintf("color:%v", tex.GetPixel(xpos, ypos)))
}

func scrollCallback(g *engine.Object, xoff float64, yoff float64) {

}
