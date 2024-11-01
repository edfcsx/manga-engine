package manga

import (
	"manga_engine/vector"
)

const KeyboardMoveComponentType = "KeyboardMoveComponent"

type KeyboardMoveComponent struct {
	Entity    *Entity
	upKeys    []string
	downKeys  []string
	leftKeys  []string
	rightKeys []string
}

func makeKeyboardMoveComponent(entity *Entity) *KeyboardMoveComponent {
	return &KeyboardMoveComponent{
		Entity:    entity,
		upKeys:    make([]string, 0),
		downKeys:  make([]string, 0),
		leftKeys:  make([]string, 0),
		rightKeys: make([]string, 0),
	}
}

func (k *KeyboardMoveComponent) SetKeys(upKeys []string, downKeys []string, leftKeys []string, rightKeys []string) {
	k.upKeys = upKeys
	k.downKeys = downKeys
	k.leftKeys = leftKeys
	k.rightKeys = rightKeys
}

func (k *KeyboardMoveComponent) Initialize() {
	transform := GetTransformComponent(k.Entity)

	if transform == nil {
		return
	}

	transform.update = func(t *TransformComponent) {
		directions := vector.MakeVec2[float64](0.0, 0.0)

		if Engine.Keyboard.IsAnyKeyPressed(k.upKeys) && Engine.Keyboard.IsAnyKeyPressed(k.downKeys) {
			directions.Y = 0.0
		} else {
			if Engine.Keyboard.IsAnyKeyPressed(k.upKeys) {
				directions.Y = -1.0
			} else if Engine.Keyboard.IsAnyKeyPressed(k.downKeys) {
				directions.Y = 1.0
			}
		}

		if Engine.Keyboard.IsAnyKeyPressed(k.leftKeys) && Engine.Keyboard.IsAnyKeyPressed(k.rightKeys) {
			directions.X = 0.0
		} else {
			if Engine.Keyboard.IsAnyKeyPressed(k.leftKeys) {
				directions.X = -1.0
			} else if Engine.Keyboard.IsAnyKeyPressed(k.rightKeys) {
				directions.X = 1.0
			}
		}

		if directions.X != 0.0 || directions.Y != 0.0 {
			velocity := vector.MakeVec2[float64](
				t.velocity.X*directions.X,
				t.velocity.Y*directions.Y,
			)

			t.position.X = t.position.X + (velocity.X * Engine.deltaTime)
			t.position.Y = t.position.Y + (velocity.Y * Engine.deltaTime)
		}
	}
}

func (k *KeyboardMoveComponent) Update() {}

func (k *KeyboardMoveComponent) Render() {}

func GetKeyboardMoveComponent(entity *Entity) *KeyboardMoveComponent {
	component := entity.GetComponent(KeyboardMoveComponentType)

	if component == nil {
		return nil
	}

	return component.(*KeyboardMoveComponent)
}
