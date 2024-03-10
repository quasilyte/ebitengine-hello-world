package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/quasilyte/ebitengine-hello-world/internal/assets"
	"github.com/quasilyte/ebitengine-hello-world/internal/controls"
	"github.com/quasilyte/ebitengine-hello-world/internal/game"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

type Player struct {
	pos gmath.Vec // {X, Y}
	img *ebiten.Image
}

type myGame struct {
	inputSystem input.System
	ctx         *game.Context

	player *Player
}

func main() {
	ctx := &game.Context{
		Loader:       createLoader(),
		WindowWidth:  320,
		WindowHeight: 240,
	}
	g := &myGame{
		ctx: ctx,
	}
	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	ctx.Input = g.inputSystem.NewHandler(0, controls.DefaultKeymap)

	ebiten.SetWindowSize(g.ctx.WindowWidth, g.ctx.WindowHeight)
	ebiten.SetWindowTitle("Ebitengine Quest")

	assets.RegisterResources(ctx.Loader)

	g.init()

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

func (g *myGame) Update() error {
	g.inputSystem.Update()

	speed := 64.0 * (1.0 / 60)
	var v gmath.Vec
	if g.ctx.Input.ActionIsPressed(controls.ActionMoveRight) {
		v.X += speed
	}
	if g.ctx.Input.ActionIsPressed(controls.ActionMoveDown) {
		v.Y += speed
	}
	if g.ctx.Input.ActionIsPressed(controls.ActionMoveLeft) {
		v.X -= speed
	}
	if g.ctx.Input.ActionIsPressed(controls.ActionMoveUp) {
		v.Y -= speed
	}
	g.player.pos = g.player.pos.Add(v)

	return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {
	var options ebiten.DrawImageOptions
	options.GeoM.Translate(g.player.pos.X, g.player.pos.Y)
	screen.DrawImage(g.player.img, &options)
}

func (g *myGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ctx.WindowWidth, g.ctx.WindowHeight
}

func (g *myGame) init() {
	gopher := g.ctx.Loader.LoadImage(assets.ImageGopher).Data
	g.player = &Player{img: gopher}
}

func createLoader() *resource.Loader {
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAsset
	return loader
}
