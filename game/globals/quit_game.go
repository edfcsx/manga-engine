package globals

import (
	"manga_engine/manga"
)

type QuitGame struct{}

func (q QuitGame) Handler() {
	if manga.Engine.Keyboard.IsKeyPressed("Escape") {
		manga.Engine.Quit()
	}
}
