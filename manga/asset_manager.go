package manga

import (
	"encoding/json"
	"fmt"
	"os"
)

var textures map[string]*Texture = make(map[string]*Texture)

type AssetManager struct{}

type assetManagerJSON struct {
	Textures []assetManagerTextureJSON `json:"textures"`
}

type assetManagerTextureJSON struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

func makeAssetManager() AssetManager {
	return AssetManager{}
}

func (a *AssetManager) Clear() {
	textures = make(map[string]*Texture)
}

func (a *AssetManager) AddTexture(id string, file string) {
	textures[id] = Engine.TextureManager.LoadTexture(file)
}

func (a *AssetManager) GetTexture(id string) *Texture {
	if texture, ok := textures[id]; ok {
		return texture
	}

	return nil
}

func (a *AssetManager) LoadFromJSON(path string) {
	b, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("error opening assets file")
		panic(err)
	}

	var assets assetManagerJSON
	err = json.Unmarshal(b, &assets)

	if err != nil {
		fmt.Println("error parsing assets file")
		panic(err)
	}

	for _, texture := range assets.Textures {
		a.AddTexture(texture.Id, texture.Path)
	}
}
