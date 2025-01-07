package manga

import "github.com/veandco/go-sdl2/sdl"

type TileSet struct {
	TileWidth  int32
	TileHeight int32
	Texture    *Texture
}

type TileSetE struct {
	tileSets map[string]*TileSet
}

func makeTileSetE() TileSetE {
	return TileSetE{
		tileSets: make(map[string]*TileSet),
	}
}

func (t *TileSetE) AddTileSet(label string, tileWidth int32, tileHeight int32, texture *Texture) {
	t.tileSets[label] = &TileSet{
		TileWidth:  tileWidth,
		TileHeight: tileHeight,
		Texture:    texture,
	}
}

func (t *TileSetE) GetTileSet(label string) *TileSet {
	if tileSet, ok := t.tileSets[label]; ok {
		return tileSet
	}

	return nil
}

func (t *TileSet) GetTileRect(tileID int32) *sdl.Rect {
	columns := t.Texture.Width / t.TileWidth
	x := (tileID % columns) * t.TileWidth
	y := (tileID / columns) * t.TileHeight

	return &sdl.Rect{
		X: x,
		Y: y,
		W: t.TileWidth,
		H: t.TileHeight,
	}
}
