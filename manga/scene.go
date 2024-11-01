package manga

type Scene interface {
	Initialize()
	Update()
	Render()
}
