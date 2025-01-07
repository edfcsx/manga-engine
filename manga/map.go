package manga

import (
	"encoding/json"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"manga_engine/vector"
	"os"
)

type Map struct {
	Dimensions vector.Vec2[int32]
	layers     []MapLayer
}

type MapLayer struct {
	tiles []Tile
}

type Tile struct {
	texture *Texture
	src     *sdl.Rect
	dst     *sdl.Rect
}

type MapData struct {
	CompressionLevel int              `json:"compressionlevel"`
	Height           int              `json:"height"`
	Width            int              `json:"width"`
	Infinite         bool             `json:"infinite"`
	Layers           []MapDataLayer   `json:"layers"`
	NextLayerID      int              `json:"nextlayerid"`
	NextObjectID     int              `json:"nextobjectid"`
	Orientation      string           `json:"orientation"`
	RenderOrder      string           `json:"renderorder"`
	TiledVersion     string           `json:"tiledversion"`
	TileHeight       int              `json:"tileheight"`
	TileWidth        int              `json:"tilewidth"`
	Type             string           `json:"type"`
	Version          string           `json:"version"`
	TileSets         []MapDataTileSet `json:"tilesets"`
}

type MapDataLayer struct {
	Name    string  `json:"name"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
	Width   int     `json:"width"`
	Height  int     `json:"height"`
	Data    []int   `json:"data"`
	Type    string  `json:"type"`
	Visible bool    `json:"visible"`
	Opacity float32 `json:"opacity"`
}

type MapDataTileSet struct {
	FirstGid int    `json:"firstgid"`
	Source   string `json:"source"`
	Name     string `json:"name"`
}

func CreateMapFromJSON(path string) (*Map, error) {
	bytes, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("error opening map file: %v", err)
	}

	var mapData MapData
	err = json.Unmarshal(bytes, &mapData)

	if err != nil {
		return nil, fmt.Errorf("error parsing map file: %v", err)
	}

	world := &Map{
		Dimensions: vector.MakeVec2[int32](int32(mapData.Width), int32(mapData.Height)),
		layers:     make([]MapLayer, len(mapData.Layers)),
	}

	for layerIdx, layer := range mapData.Layers {
		tiles := make([]Tile, 0)

		for index, tileID := range layer.Data {
			if tileID <= 0 {
				continue
			}

			tileID -= 1

			tileSet := getTileSetFromGID(tileID, mapData.TileSets)

			if tileSet == nil {
				continue
			}

			tile := Tile{
				texture: tileSet.Texture,
				src:     &sdl.Rect{X: 0, Y: 0, W: 12, H: 12},
				dst:     &sdl.Rect{X: 0, Y: 0, W: 12, H: 12},
			}

			//rows := tileSet.Texture.Height / tileSet.TileHeight
			//columns := tileSet.Texture.Width / tileSet.TileWidth

			tile.src.X = int32(tileID%32) * 12
			tile.src.Y = int32(tileID/32) * 12

			tile.dst.X = int32(int32(index)%int32(mapData.Width)) * 12
			tile.dst.Y = int32(int32(index)/int32(mapData.Width)) * 12

			//fmt.Printf("TileID: %d, Index: %d, Src: (%d, %d), Dst: (%d, %d)\n",
			//	tileID, index, tile.src.X, tile.src.Y, tile.dst.X, tile.dst.Y)

			tiles = append(tiles, tile)
		}

		world.layers[layerIdx] = MapLayer{
			tiles: tiles,
		}

	}

	return world, nil
}

func (m *Map) Draw() {
	for _, layer := range m.layers {
		for _, tile := range layer.tiles {
			Engine.TextureManager.Draw(tile.texture, tile.src, tile.dst, sdl.FLIP_NONE)
		}
	}
}

func (m *Map) DrawLayer(layer int) {
	for _, tile := range m.layers[layer].tiles {
		Engine.TextureManager.Draw(tile.texture, tile.src, tile.dst, sdl.FLIP_NONE)
	}
}

func getTileSetFromGID(gid int, tileSets []MapDataTileSet) *TileSet {
	for idx, tileSet := range tileSets {
		if idx+1 < len(tileSets) {
			nextTileSet := tileSets[idx+1]

			if gid < nextTileSet.FirstGid {
				return Engine.TileSet.GetTileSet(tileSet.Name)
			}
		} else {
			if gid >= tileSet.FirstGid {
				return Engine.TileSet.GetTileSet(tileSet.Name)
			}
		}
	}

	return nil
}
