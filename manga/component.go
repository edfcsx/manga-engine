package manga

type Component interface {
	Initialize()
	Update()
	Render()
}

type ComponentE struct{}
