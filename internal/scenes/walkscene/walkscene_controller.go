package walkscene

import (
	"github.com/quasilyte/ebitengine-hello-world/internal/game"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

type scene = gscene.Scene[*WalksceneController]

type WalksceneController struct {
	ctx *game.Context

	scene *gscene.RootScene[*WalksceneController]
}

func NewWalksceneController(ctx *game.Context) *WalksceneController {
	return &WalksceneController{ctx: ctx}
}

// Тип аргумента scene определяет какой интерфейс должен реализовывать контроллер.
// В данном случае, WalksceneController должен реализовать controllerAccessor.
func (c *WalksceneController) Init(s *gscene.RootScene[*WalksceneController]) {
	c.scene = s

	g := newGopherNode(gmath.Vec{X: 64, Y: 64})
	s.AddObject(g)
}

func (c *WalksceneController) Update(delta float64) {
}
