package walkscene

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gsignal"
)

type pickupNode struct {
	pos      gmath.Vec
	rect     *graphics.Rect
	scene    *scene
	score    int
	disposed bool

	EventDestroyed gsignal.Event[int]
}

func newPickupNode(pos gmath.Vec) *pickupNode {
	return &pickupNode{pos: pos}
}

func (n *pickupNode) Init(s *scene) {
	n.scene = s
	ctx := s.Controller().ctx

	n.score = ctx.Rand.IntRange(5, 10)

	n.rect = ctx.NewRect(16, 16)
	n.rect.Pos.Base = &n.pos
	n.rect.SetFillColorScale(graphics.ColorScaleFromRGBA(200, 200, 0, 255))
	s.AddGraphics(n.rect)
}

func (n *pickupNode) IsDisposed() bool {
	return n.disposed
}

func (n *pickupNode) Update(delta float64) {
	g := n.scene.Controller().state.gopher
	if g.pos.DistanceTo(n.pos) < 16 {
		n.pickUp()
	}
}

func (n *pickupNode) pickUp() {
	n.EventDestroyed.Emit(n.score)
	n.dispose()
}

func (n *pickupNode) dispose() {
	n.rect.Dispose()
	n.disposed = true
}
