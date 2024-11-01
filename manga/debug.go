package manga

import (
	"time"
)

type Debug struct {
	enabled bool
	FPS     FPSCounter
}

func makeDebug() Debug {
	return Debug{
		enabled: false,
		FPS:     FPSCounter{},
	}
}

func (d *Debug) Enable() {
	d.enabled = true
	d.FPS.initialize()
}

func (d *Debug) update() {
	if d.enabled {
		d.FPS.update()
	}
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
