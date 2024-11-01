package main

import (
	"manga_engine/game"
	"manga_engine/game/globals"
	"manga_engine/manga"
)

func main() {
	manga.Engine.Window.Size(800, 600)
	manga.Engine.Window.Position(manga.WINDOW_CENTERED, manga.WINDOW_CENTERED)

	// enable debug
	manga.Engine.Debug.Enable()

	// adding globals scripts
	manga.Engine.AddGlobalScript(globals.QuitGame{})

	manga.Start(&game.Home{})
}
