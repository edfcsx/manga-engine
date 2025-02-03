package game

import (
	"manga_engine/manga"
)

type Home struct {
	entities manga.EntityManager
	world    *manga.Map
}

func (h *Home) Initialize() {
	// loading assets
	manga.Engine.AssetManager.LoadFromJSON("game/assets.json")

	// adding tileset
	islandTexture := manga.Engine.AssetManager.GetTexture("island_tileset")
	manga.Engine.TileSet.AddTileSet("island", 12, 12, islandTexture)

	// adding map
	world, err := manga.CreateMapFromJSON("assets/maps/town2.json")
	if err != nil {
		panic(err)
	}

	h.world = world

	// adding entities
	h.entities = manga.Engine.EntityManager.Make()
	h.entities.AddEntity(MakePlayer().GetEntity())

	/* radar entity */
	radar := h.entities.CreateEntity("radar")
	radar.CreateTransformE(720.0, 20.0, 0.0, 0.0, 64, 64, 1)

	radarSpr := radar.CreateSprite("radar")
	radarSpr.AddAnimation("idle", 0, 8, 200, true)
	radarSpr.PlayAnimation("idle")
}

func (h *Home) Update() {
	h.entities.Update()
	// show fps
	//fmt.Printf("FPS: %f\n", manga.Engine.Debug.FPS.GetFPS())
}

func (h *Home) Render() {
	h.world.Draw()
	h.entities.Render()
}
