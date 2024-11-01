package manga

func makeEntity(label string) *Entity {
	return &Entity{
		Label:      label,
		IsActive:   true,
		components: make(map[string]Component),
	}
}

type Entity struct {
	Label      string
	IsActive   bool
	components map[string]Component
}

func (e *Entity) AddComponent(componentType string, c Component) {
	e.components[componentType] = c
	c.Initialize()
}

func (e *Entity) GetComponent(componentType string) Component {
	return e.components[componentType]
}

func (e *Entity) Update() {
	for _, c := range e.components {
		c.Update()
	}
}

func (e *Entity) Render() {
	for _, c := range e.components {
		c.Render()
	}
}

func (e *Entity) CreateTransform() *TransformComponent {
	transform := makeTransformComponent(e)
	e.AddComponent(TransformComponentType, transform)
	return transform
}

func (e *Entity) CreateTransformE(posX float64, posY float64, velX float64, velY float64, sizeW int32, sizeH int32, scale int32) *TransformComponent {
	transform := makeTransformComponent(e)
	transform.Position(posX, posY)
	transform.Velocity(velX, velY)
	transform.Size(sizeW, sizeH)
	transform.Scale(scale)

	e.AddComponent(TransformComponentType, transform)

	return transform
}

func (e *Entity) CreateSprite(texture string) *SpriteComponent {
	sprite := makeSpriteComponent(e, texture)
	e.AddComponent(SpriteComponentType, sprite)
	return sprite
}

func (e *Entity) CreateScript(initialize func(*Entity), update func(*Entity), render func(*Entity)) *ScriptComponent {
	script := makeScriptComponent(e, initialize, update, render)
	e.AddComponent(ScriptComponentType, script)
	return script
}

func (e *Entity) CreateKeyboardMove() *KeyboardMoveComponent {
	keyboardMove := makeKeyboardMoveComponent(e)
	e.AddComponent(KeyboardMoveComponentType, keyboardMove)
	return keyboardMove
}
