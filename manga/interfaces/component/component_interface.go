package componentI

type Component interface {
	GetType() int32
	Initialize()
	Update(deltaTime float64)
	Render()
}
