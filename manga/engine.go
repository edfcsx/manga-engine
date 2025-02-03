package manga

import (
	"github.com/veandco/go-sdl2/sdl"
)

var globalScripts = make([]GlobalScript, 0)

type engine struct {
	Window         Window
	gameWindow     *sdl.Window
	renderer       *sdl.Renderer
	running        bool
	deltaTime      float64
	ticksLastFrame uint64
	Graphics       Graphics
	currentScene   Scene
	EntityManager  EntityManagerE
	TextureManager textureManager
	AssetManager   AssetManager
	Debug          Debug
	Keyboard       Keyboard
	TileSet        TileSetE
	systemManager  SystemManager
}

var Engine = &engine{
	Window:         makeWindow(),
	Graphics:       makeGraphics(),
	gameWindow:     nil,
	renderer:       nil,
	running:        false,
	deltaTime:      0.0,
	ticksLastFrame: 0,
	currentScene:   nil,
	EntityManager:  EntityManagerE{},
	TextureManager: makeTextureManager(),
	AssetManager:   makeAssetManager(),
	Debug:          makeDebug(),
	Keyboard:       makeKeyboard(),
	TileSet:        makeTileSetE(),
	systemManager:  makeSystemManager(),
}

func (e *engine) AddGlobalScript(script GlobalScript) {
	globalScripts = append(globalScripts, script)
}

func Stop() {
	Engine.running = false
}

func Start(scene Scene) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	defer destroy()

	window, err := sdl.CreateWindow(
		Engine.Window.title,
		Engine.Window.pos.X,
		Engine.Window.pos.Y,
		Engine.Window.size.X,
		Engine.Window.size.Y,
		sdl.WINDOW_SHOWN,
	)

	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, 0)

	if err != nil {
		panic(err)
	}

	Engine.gameWindow = window
	Engine.renderer = renderer

	Engine.running = true

	Engine.systemManager.AddSystem(ColliderSystemID, MakeColliderSystem())

	Engine.currentScene = scene
	Engine.currentScene.Initialize()

	mainLoop()
}

func executeGlobalScripts() {
	for _, script := range globalScripts {
		script.Handler()
	}
}

func processEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			Engine.running = false
			break
		case *sdl.KeyboardEvent:
			keyEvent := event.(*sdl.KeyboardEvent)

			if keyEvent.Type == sdl.KEYDOWN {
				registerKeyPressed(int(keyEvent.Keysym.Sym))
				break
			} else if keyEvent.Type == sdl.KEYUP {
				registerKeyReleased(int(keyEvent.Keysym.Sym))
				break
			}
		}
	}
}

func update() {
	currentTicks := sdl.GetTicks64()
	timeElapsed := currentTicks - Engine.ticksLastFrame

	if timeElapsed > uint64(Engine.Graphics.frameTargetTime) {
		timeElapsed = uint64(Engine.Graphics.frameTargetTime)
	}

	timeToWait := uint64(Engine.Graphics.frameTargetTime) - timeElapsed

	if timeToWait > 0 && timeToWait <= 1000 {
		sdl.Delay(uint32(timeToWait))
		currentTicks = sdl.GetTicks64()
		timeElapsed = currentTicks - Engine.ticksLastFrame
	}

	Engine.deltaTime = float64(timeElapsed) / 1000.0

	if Engine.deltaTime > 0.05 {
		Engine.deltaTime = 0.05
	}

	Engine.ticksLastFrame = sdl.GetTicks64()
	Engine.currentScene.Update()
}

func render() {
	err := Engine.renderer.SetDrawColor(0, 0, 0, 255)

	if err != nil {
		return
	}

	err = Engine.renderer.Clear()

	Engine.currentScene.Render()

	if err != nil {
		return
	}

	Engine.renderer.Present()
}

func destroy() {
	if Engine.renderer != nil {
		err := Engine.renderer.Destroy()

		if err != nil {
			panic(err)
		}
	}

	if Engine.gameWindow != nil {
		err := Engine.gameWindow.Destroy()

		if err != nil {
			panic(err)
		}
	}

	sdl.Quit()
}

func (e *engine) Quit() {
	e.running = false
}
