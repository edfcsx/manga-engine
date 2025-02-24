package componentI

import "manga_engine/vector"

type TransformComponent interface {
	Component
	Position(x, y int32)
	Velocity(x, y int32)
	Size(x, y int32)
	Scale(int32)
	GetPosition() vector.Vec2[int32]
	GetVelocity() vector.Vec2[int32]
	GetSize() vector.Vec2[int32]
	GetScale() int32
}
