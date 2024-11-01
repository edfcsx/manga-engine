package manga

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type textureManager struct{}

func makeTextureManager() textureManager {
	return textureManager{}
}

func (t *textureManager) LoadTexture(filename string) *Texture {
	surface, err := img.Load(filename)

	if err != nil {
		panic(err)
	}

	texture, err := Engine.renderer.CreateTextureFromSurface(surface)

	if err != nil {
		panic(err)
	}

	engineTexture := &Texture{
		Width:  surface.W,
		Height: surface.H,
		source: texture,
	}

	surface.Free()
	return engineTexture
}

func (t *textureManager) Draw(texture *Texture, src *sdl.Rect, dst *sdl.Rect, flip sdl.RendererFlip) {
	err := Engine.renderer.CopyEx(texture.source, src, dst, 0, nil, flip)

	if err != nil {
		return
	}
}
