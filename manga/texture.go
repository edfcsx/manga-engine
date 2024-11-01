package manga

import "github.com/veandco/go-sdl2/sdl"

type Texture struct {
	Width  int32
	Height int32
	source *sdl.Texture
}
