package manga

type System interface {
	Initialize()
	Update()
}

type SystemManager struct {
	systems []System
}

func makeSystemManager() SystemManager {
	return SystemManager{
		systems: make([]System, 0),
	}
}

func (s *SystemManager) AddSystem(system System) {
	s.systems = append(s.systems, system)
	system.Initialize()
}

func (s *SystemManager) Update() {
	for _, system := range s.systems {
		system.Update()
	}
}
