//go:build !debug
// +build !debug

package manga

func mainLoop() {
	for Engine.running {
		executeGlobalScripts()
		processEvents()
		update()
		Engine.systemManager.Update()
		render()
	}
}
