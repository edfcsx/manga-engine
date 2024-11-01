package manga

import (
	"manga_engine/vector"
)

const TransformComponentType = "TransformComponent"

type TransformComponent struct {
	Entity   *Entity
	position vector.Vec2[float64]
	velocity vector.Vec2[float64]
	size     vector.Vec2[int32]
	scale    int32
	update   func(component *TransformComponent)
}

func makeTransformComponent(entity *Entity) *TransformComponent {
	return &TransformComponent{
		Entity:   entity,
		position: vector.MakeVec2[float64](0.0, 0.0),
		velocity: vector.MakeVec2[float64](0, 0),
		size:     vector.MakeVec2[int32](0, 0),
		scale:    1,
		update:   defaultTransformUpdate,
	}
}

func (t *TransformComponent) Initialize() {}

func (t *TransformComponent) Update() {
	t.update(t)
}

func defaultTransformUpdate(t *TransformComponent) {
	t.position.X = t.position.X + (t.velocity.X * Engine.deltaTime)
	t.position.Y = t.position.Y + (t.velocity.Y * Engine.deltaTime)
}

func (t *TransformComponent) Render() {}

func GetTransformComponent(entity *Entity) *TransformComponent {
	component := entity.GetComponent(TransformComponentType)

	if component == nil {
		return nil
	}

	return component.(*TransformComponent)
}

func (t *TransformComponent) Position(x float64, y float64) {
	t.position.X = x
	t.position.Y = y
}

func (t *TransformComponent) Velocity(x float64, y float64) {
	t.velocity.X = x
	t.velocity.Y = y
}

func (t *TransformComponent) Size(w int32, h int32) {
	t.size.X = w
	t.size.Y = h
}

func (t *TransformComponent) Scale(s int32) {
	t.scale = s
}

func (t *TransformComponent) Transform(posX float64, posY float64, velX float64, velY float64, sizeW int32, sizeH int32, scale int32) {
	t.Position(posX, posY)
	t.Velocity(velX, velY)
	t.Size(sizeW, sizeH)
	t.Scale(scale)
}

func (t *TransformComponent) GetPosition() vector.Vec2[float64] {
	return t.position
}

func (t *TransformComponent) GetVelocity() vector.Vec2[float64] {
	return t.velocity
}

func (t *TransformComponent) GetSize() vector.Vec2[int32] {
	return t.size
}

func (t *TransformComponent) GetScale() int32 {
	return t.scale
}
