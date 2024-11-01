package manga

type EntityManagerE struct{}

func (e *EntityManagerE) Make() EntityManager {
	return EntityManager{
		entities: make(map[string]*Entity),
	}
}

type EntityManager struct {
	entities map[string]*Entity
}

func (e *EntityManager) CreateEntity(label string) *Entity {
	entity := makeEntity(label)
	e.AddEntity(entity)
	return entity
}

func (e *EntityManager) AddEntity(entity *Entity) {
	e.entities[entity.Label] = entity
}

func (e *EntityManager) GetEntity(label string) *Entity {
	if entity, ok := e.entities[label]; ok {
		return entity
	}

	return nil
}

func (e *EntityManager) RemoveEntity(label string) {
	delete(e.entities, label)
}

func (e *EntityManager) Update() {
	for _, entity := range e.entities {
		entity.Update()
	}
}

func (e *EntityManager) Render() {
	for _, entity := range e.entities {
		entity.Render()
	}
}

func (e *EntityManager) Clear() {
	e.entities = make(map[string]*Entity)
}
