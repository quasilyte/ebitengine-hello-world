package game

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

type Context struct {
	Input  *input.Handler
	Loader *resource.Loader
	Rand   gmath.Rand

	WindowWidth  int
	WindowHeight int

	graphicsCache *graphics.Cache

	scene gscene.GameRunner
}

func ChangeScene[ControllerAccessor any](ctx *Context, c gscene.Controller[ControllerAccessor]) {
	s := gscene.NewRootScene[ControllerAccessor](c)
	ctx.scene = s
}

func NewContext() *Context {
	return &Context{
		graphicsCache: graphics.NewCache(),
	}
}

func (ctx *Context) NewRect(width, height float64) *graphics.Rect {
	return graphics.NewRect(ctx.graphicsCache, width, height)
}

func (ctx *Context) NewLabel(id resource.FontID) *graphics.Label {
	fnt := ctx.Loader.LoadFont(id)
	return graphics.NewLabel(ctx.graphicsCache, fnt.Face)
}

func (ctx *Context) NewSprite(id resource.ImageID) *graphics.Sprite {
	s := graphics.NewSprite(ctx.graphicsCache)
	if id == 0 {
		return s
	}
	img := ctx.Loader.LoadImage(id)
	s.SetImage(img.Data)
	if img.DefaultFrameWidth != 0 || img.DefaultFrameHeight != 0 {
		w, h := s.GetFrameSize()
		if img.DefaultFrameWidth != 0 {
			w = int(img.DefaultFrameWidth)
		}
		if img.DefaultFrameHeight != 0 {
			h = int(img.DefaultFrameHeight)
		}
		s.SetFrameSize(w, h)
	}
	return s
}

func (ctx *Context) CurrentScene() gscene.GameRunner {
	return ctx.scene
}
