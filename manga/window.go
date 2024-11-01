package manga

import (
	"github.com/veandco/go-sdl2/sdl"
	"manga_engine/vector"
)

type Window struct {
	title string
	size  vector.Vec2[int32]
	pos   vector.Vec2[int32]
}

func makeWindow() Window {
	return Window{
		title: "",
		size:  vector.MakeVec2[int32](0, 0),
		pos:   vector.MakeVec2[int32](0, 0),
	}
}

func (e *Window) Title(title string) {
	e.title = title

	if Engine.gameWindow != nil {
		Engine.gameWindow.SetTitle(title)
	}
}

func (e *Window) Size(w int32, h int32) {
	e.size = vector.MakeVec2[int32](w, h)
}

func (e *Window) Position(x int32, y int32) {
	e.pos = vector.MakeVec2[int32](x, y)
}

const WINDOW_CENTERED = sdl.WINDOWPOS_CENTERED
