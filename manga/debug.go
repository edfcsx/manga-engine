package manga

import (
	"time"
)

type Debug struct {
	FPS                FPSCounter
	ShowCollisionBoxes bool
}

func makeDebug() Debug {
	return Debug{
		FPS: FPSCounter{},
	}
}

func (d *Debug) Enable() {
	d.FPS.initialize()
}

func (d *Debug) ShowCollisions(state bool) {
	d.ShowCollisionBoxes = state
}

func (d *Debug) update() {
	d.FPS.update()
}

type FPSCounter struct {
	lastTime time.Time
	frames   int
	fps      float64
}

func (f *FPSCounter) initialize() {
	f.lastTime = time.Now()
	f.frames = 0
	f.fps = 0.0
}

func (f *FPSCounter) update() {
	now := time.Now()
	elapsed := now.Sub(f.lastTime).Seconds()
	f.frames++

	if elapsed >= 1.0 {
		f.fps = float64(f.frames) / elapsed
		f.frames = 0
		f.lastTime = now
	}
}

func (f *FPSCounter) GetFPS() float64 {
	return f.fps
}
