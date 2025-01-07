package main

import (
	"manga_engine/game"
	"manga_engine/game/globals"
	"manga_engine/manga"
)

func main() {
	manga.Engine.Window.Size(800, 600)
	manga.Engine.Window.Position(manga.WINDOW_CENTERED, manga.WINDOW_CENTERED)
	manga.Engine.Window.Title("Manga Engine")

	// enable debug
	manga.Engine.Debug.Enable()
	manga.Engine.Debug.ShowCollisions(true)

	// adding globals scripts
	manga.Engine.AddGlobalScript(globals.QuitGame{})

	manga.Start(&game.Home{})
}
