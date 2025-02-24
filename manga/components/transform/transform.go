package transformComponent

import (
	entityI "manga_engine/manga/interfaces/entity"
	"manga_engine/vector"
)

type Transform struct {
	Entity   *entityI.EntityI
	position vector.Vec2[int32]
	velocity vector.Vec2[int32]
	size     vector.Vec2[int32]
	scale    int32
}

func Make(e *entityI.EntityI) *Transform {
	return &Transform{
		Entity: e,
	}
}

func (t *Transform) GetType() int32 { return 0 }

func (t *Transform) Initialize() {}

func (t *Transform) Update(deltaTime float64) {}

func (t *Transform) Render() {}

func (t *Transform) Position(x int32, y int32) {
	t.position.X = x
	t.position.Y = y
}

func (t *Transform) Velocity(x int32, y int32) {
	t.velocity.X = x
	t.velocity.Y = y
}

func (t *Transform) Size(x int32, y int32) {
	t.size.X = x
	t.size.Y = y
}

func (t *Transform) Scale(s int32) {
	t.scale = s
}

func (t *Transform) GetPosition() vector.Vec2[int32] {
	return t.position
}

func (t *Transform) GetVelocity() vector.Vec2[int32] {
	return t.velocity
}

func (t *Transform) GetSize() vector.Vec2[int32] {
	return t.size
}

func (t *Transform) GetScale() int32 {
	return t.scale
}
