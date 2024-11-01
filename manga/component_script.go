package manga

const ScriptComponentType = "ScriptComponent"

type ScriptComponent struct {
	Entity     *Entity
	initialize func(*Entity)
	update     func(*Entity)
	render     func(*Entity)
}

func makeScriptComponent(entity *Entity, initialize func(*Entity), update func(*Entity), render func(*Entity)) *ScriptComponent {
	return &ScriptComponent{
		Entity:     entity,
		initialize: initialize,
		update:     update,
		render:     render,
	}
}

func (s *ScriptComponent) Initialize() {
	if s.initialize != nil {
		s.initialize(s.Entity)
	}
}

func (s *ScriptComponent) Update() {
	if s.update != nil {
		s.update(s.Entity)
	}
}

func (s *ScriptComponent) Render() {
	if s.render != nil {
		s.render(s.Entity)
	}
}

func GetScriptComponent(entity *Entity) *ScriptComponent {
	component := entity.GetComponent(ScriptComponentType)

	if component == nil {
		return nil
	}

	return component.(*ScriptComponent)
}
