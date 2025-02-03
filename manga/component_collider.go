package manga

import "fmt"

const ColliderComponentType = "ColliderComponent"

type ColliderComponent struct {
	Entity      *Entity
	OnCollision func(*Entity)
	Shape       ColliderShape
	transform   *TransformComponent
}

func makeColliderComponent(entity *Entity, OnCollision func(*Entity), shape ColliderShape) *ColliderComponent {
	if shape == nil {
		panic(fmt.Sprintf("ColliderComponent: shape is nil for entity %s", entity.Label))
	} else if entity.GetComponent(TransformComponentType) == nil {
		panic(fmt.Sprintf("ColliderComponent: entity %s does not have TransformComponent", entity.Label))
	}

	return &ColliderComponent{
		Entity:      entity,
		OnCollision: OnCollision,
		Shape:       shape,
		transform:   GetTransformComponent(entity),
	}
}

func (c *ColliderComponent) Initialize() {
	Engine.systemManager.systems[ColliderSystemID].(*ColliderSystem).Register(ColliderMoving, c.Shape, c.OnCollision)
}

func (c *ColliderComponent) Update() {
	transform := GetTransformComponent(c.Entity)
	c.Shape.MoveTo(int32(transform.position.X), int32(transform.position.Y))
}

func (c *ColliderComponent) Render() {
	if Engine.Debug.ShowCollisionBoxes {
		c.Shape.Render(c.transform)
	}
}

func GetColliderComponent(entity *Entity) *ColliderComponent {
	component := entity.GetComponent(ColliderComponentType)

	if component == nil {
		return nil
	}

	return component.(*ColliderComponent)
}
