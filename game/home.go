package game

import (
	"fmt"
	"manga_engine/manga"
)

type Home struct {
	entities manga.EntityManager
}

func (h *Home) Initialize() {
	manga.Engine.AssetManager.LoadFromJSON("game/assets.json")
	h.entities = manga.Engine.EntityManager.Make()

	chopper := h.entities.CreateEntity("chopper")

	chopper.CreateScript(nil, chopperScriptUpdate, nil)
	chopper.CreateTransformE(400.0, 200.0, 200.0, 200.0, 32, 32, 3)
	chopper.CreateKeyboardMove().SetKeys(moveUpKeys, moveDownKeys, moveLeftKeys, moveRightKeys)

	chopperSpr := chopper.CreateSprite("chopper")
	chopperSpr.AddAnimation("down", 0, 2, 80, false)
	chopperSpr.AddAnimation("right", 1, 2, 80, false)
	chopperSpr.AddAnimation("left", 2, 2, 80, false)
	chopperSpr.AddAnimation("up", 3, 2, 80, false)
	chopperSpr.PlayAnimation("down")

	/* radar entity */
	radar := h.entities.CreateEntity("radar")
	radar.CreateTransformE(720.0, 20.0, 0.0, 0.0, 64, 64, 1)

	radarSpr := radar.CreateSprite("radar")
	radarSpr.AddAnimation("idle", 0, 8, 200, true)
	radarSpr.PlayAnimation("idle")
}

func chopperScriptUpdate(e *manga.Entity) {
	spr := manga.GetSpriteComponent(e)

	if manga.Engine.Keyboard.IsAnyKeyPressed(moveUpKeys) {
		spr.PlayAnimation("up")
	} else if manga.Engine.Keyboard.IsAnyKeyPressed(moveDownKeys) {
		spr.PlayAnimation("down")
	} else if manga.Engine.Keyboard.IsAnyKeyPressed(moveLeftKeys) {
		spr.PlayAnimation("left")
	} else if manga.Engine.Keyboard.IsAnyKeyPressed(moveRightKeys) {
		spr.PlayAnimation("right")
	}
}

func (h *Home) Update() {
	h.entities.Update()

	fmt.Println("fps: ", manga.Engine.Debug.FPS.GetFPS())
}

func (h *Home) Render() {
	h.entities.Render()
}
