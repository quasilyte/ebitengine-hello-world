package scenes

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/ebitengine-hello-world/internal/assets"
	"github.com/quasilyte/ebitengine-hello-world/internal/controls"
	"github.com/quasilyte/ebitengine-hello-world/internal/game"
	"github.com/quasilyte/ebitengine-hello-world/internal/scenes/walkscene"
	"github.com/quasilyte/gscene"
)

type SplashController struct {
	ctx *game.Context
}

func NewSplashController(ctx *game.Context) *SplashController {
	return &SplashController{ctx: ctx}
}

func (c *SplashController) Init(s *gscene.SimpleRootScene) {
	l := c.ctx.NewLabel(assets.FontBig)
	l.SetAlignHorizontal(graphics.AlignHorizontalCenter)
	l.SetAlignVertical(graphics.AlignVerticalCenter)
	l.SetSize(c.ctx.WindowWidth, c.ctx.WindowHeight)
	l.SetText("Press [Enter] to continue")
	s.AddGraphics(l)
}

func (c *SplashController) Update(delta float64) {
	if c.ctx.Input.ActionIsJustPressed(controls.ActionConfirm) {
		game.ChangeScene(c.ctx, walkscene.NewWalksceneController(c.ctx))
	}
}
