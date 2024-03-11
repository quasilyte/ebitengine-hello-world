package walkscene

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/ebitengine-hello-world/internal/assets"
	"github.com/quasilyte/ebitengine-hello-world/internal/controls"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
)

type gopherNode struct {
	input  *input.Handler
	pos    gmath.Vec
	sprite *graphics.Sprite
}

func newGopherNode(pos gmath.Vec) *gopherNode {
	return &gopherNode{
		pos: pos,
	}
}

func (g *gopherNode) Init(s *scene) {
	ctx := s.Controller().ctx

	g.input = ctx.Input

	g.sprite = ctx.NewSprite(assets.ImageGopher)
	g.sprite.Pos.Base = &g.pos
	s.AddGraphics(g.sprite)
}

func (g *gopherNode) IsDisposed() bool {
	return false
}

func (g *gopherNode) Update(delta float64) {
	speed := 64.0 * (1.0 / 60)
	var v gmath.Vec

	if g.input.ActionIsPressed(controls.ActionMoveRight) {
		v.X += speed
	}

	if g.input.ActionIsPressed(controls.ActionMoveDown) {
		v.Y += speed
	}

	if g.input.ActionIsPressed(controls.ActionMoveLeft) {
		v.X -= speed
	}

	if g.input.ActionIsPressed(controls.ActionMoveUp) {
		v.Y -= speed
	}

	g.pos = g.pos.Add(v)
}
