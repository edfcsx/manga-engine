package manga

type System interface {
	Initialize()
	Update()
}

type SystemManager struct {
	systems map[string]System
}

func makeSystemManager() SystemManager {
	return SystemManager{
		systems: make(map[string]System),
	}
}

func (s *SystemManager) AddSystem(id string, system System) {
	s.systems[id] = system
	system.Initialize()
}

func (s *SystemManager) Update() {
	for _, system := range s.systems {
		system.Update()
	}
}
