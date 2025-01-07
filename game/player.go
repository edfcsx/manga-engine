package game

import (
	"manga_engine/manga"
)

const SprAnimSpeed = 80

// Player implements custom entity
type Player struct {
	entity    *manga.Entity
	transform *manga.TransformComponent
	sprite    *manga.SpriteComponent
	collider  *manga.ColliderComponent
}

func (p *Player) GetEntity() *manga.Entity {
	return p.entity
}

func MakePlayer() *Player {
	player := &Player{
		entity:    manga.Engine.EntityManager.CreateEntity("player"),
		transform: nil,
		sprite:    nil,
		collider:  nil,
	}

	player.entity.SetSelf(player)

	// create transform
	var playerScale int32 = 3
	transform := player.entity.CreateTransformE(400.0, 200.0, 200.0, 200.0, 32, 32, 3)
	player.transform = transform

	// create sprite and animations
	sprite := player.entity.CreateSprite("chopper")
	sprite.AddAnimation("down", 0, 2, SprAnimSpeed, false)
	sprite.AddAnimation("right", 1, 2, SprAnimSpeed, false)
	sprite.AddAnimation("left", 2, 2, SprAnimSpeed, false)
	sprite.AddAnimation("up", 3, 2, SprAnimSpeed, false)
	sprite.PlayAnimation("down")
	player.sprite = sprite

	// create custom scripts for player
	player.entity.CreateScript(nil, update, nil)
	//player.entity.CreateScript(player, nil, nil, nil)
	//player.entity.CreateScript(player, nil, update, nil)

	// create keyboard move component
	player.entity.CreateKeyboardMove().SetKeys(moveUpKeys, moveDownKeys, moveLeftKeys, moveRightKeys)

	// create collider component
	collider := player.entity.CreateCollider(collisionHandler, manga.MakeCircleShape(16*playerScale))
	player.collider = collider

	return player
}

func update(e *manga.Entity) {
	self := e.Self.(*Player)

	if manga.Engine.Keyboard.IsAnyKeyPressed(moveUpKeys) {
		self.sprite.PlayAnimation("up")
	} else if manga.Engine.Keyboard.IsAnyKeyPressed(moveDownKeys) {
		self.sprite.PlayAnimation("down")
	} else if manga.Engine.Keyboard.IsAnyKeyPressed(moveLeftKeys) {
		self.sprite.PlayAnimation("left")
	} else if manga.Engine.Keyboard.IsAnyKeyPressed(moveRightKeys) {
		self.sprite.PlayAnimation("right")
	}

}

func collisionHandler(e *manga.Entity) {

}
