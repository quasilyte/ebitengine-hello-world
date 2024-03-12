package walkscene

import (
	"fmt"

	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/ebitengine-hello-world/internal/assets"
	"github.com/quasilyte/ebitengine-hello-world/internal/controls"
	"github.com/quasilyte/ebitengine-hello-world/internal/game"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

type scene = gscene.Scene[*Controller]

type Controller struct {
	ctx *game.Context

	state *sceneState

	scene *gscene.RootScene[*Controller]

	scoreLabel *graphics.Label
	score      int
}

func NewController(ctx *game.Context) *Controller {
	return &Controller{ctx: ctx}
}

func (c *Controller) Init(s *gscene.RootScene[*Controller]) {
	c.scene = s

	g := newGopherNode(gmath.Vec{X: 64, Y: 64})
	s.AddObject(g)

	c.state = &sceneState{gopher: g}

	c.scoreLabel = c.ctx.NewLabel(assets.FontNormal)
	c.scoreLabel.Pos.Offset = gmath.Vec{X: 4, Y: 4}
	s.AddGraphics(c.scoreLabel)

	c.createPickup()
	c.addScore(0)
}

func (c *Controller) createPickup() {
	p := newPickupNode(gmath.Vec{
		X: c.ctx.Rand.FloatRange(0, float64(c.ctx.WindowWidth)),
		Y: c.ctx.Rand.FloatRange(0, float64(c.ctx.WindowHeight)),
	})

	p.EventDestroyed.Connect(nil, func(score int) {
		c.addScore(score)
		c.createPickup()
	})

	c.scene.AddObject(p)
}

func (c *Controller) addScore(score int) {
	c.score += score
	c.scoreLabel.SetText(fmt.Sprintf("score: %d", c.score))
}

func (c *Controller) Update(delta float64) {
	if c.ctx.Input.ActionIsJustPressed(controls.ActionRestart) {
		game.ChangeScene(c.ctx, NewController(c.ctx))
	}
}
