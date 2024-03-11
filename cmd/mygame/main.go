package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/quasilyte/ebitengine-hello-world/internal/assets"
	"github.com/quasilyte/ebitengine-hello-world/internal/controls"
	"github.com/quasilyte/ebitengine-hello-world/internal/game"
	"github.com/quasilyte/ebitengine-hello-world/internal/scenes"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
)

type myGame struct {
	inputSystem input.System
	ctx         *game.Context
}

func main() {
	ctx := game.NewContext()
	ctx.Loader = createLoader()
	ctx.WindowWidth = 320
	ctx.WindowHeight = 240
	ctx.Rand.SetSeed(time.Now().Unix())
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

	game.ChangeScene(ctx, scenes.NewSplashController(ctx))

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

func (g *myGame) Update() error {
	g.inputSystem.Update()
	g.ctx.CurrentScene().UpdateWithDelta(1.0 / 60.0)
	return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {
	g.ctx.CurrentScene().Draw(screen)
}

func (g *myGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ctx.WindowWidth, g.ctx.WindowHeight
}

func createLoader() *resource.Loader {
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAsset
	return loader
}
