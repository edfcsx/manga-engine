package manga

import "github.com/veandco/go-sdl2/sdl"

var windowKeyMap = map[int]bool{}

func registerKeyPressed(key int) {
	windowKeyMap[key] = true
}

func registerKeyReleased(key int) {
	windowKeyMap[key] = false
}

type Keyboard struct{}

func makeKeyboard() Keyboard {
	return Keyboard{}
}

func (k *Keyboard) IsKeyPressed(key string) bool {
	code := sdl.GetKeyFromName(key)

	if windowKeyMap[int(code)] {
		return true
	}

	return false
}

func (k *Keyboard) IsAnyKeyPressed(keys []string) bool {
	for _, key := range keys {
		if k.IsKeyPressed(key) {
			return true
		}
	}

	return false
}

func (k *Keyboard) IsAllKeysPressed(keys []string) bool {
	for _, key := range keys {
		if !k.IsKeyPressed(key) {
			return false
		}
	}

	return true
}
