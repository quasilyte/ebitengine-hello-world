package game

import (
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
)

type Context struct {
	Input  *input.Handler
	Loader *resource.Loader

	WindowWidth  int
	WindowHeight int
}
