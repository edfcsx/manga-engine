package manga

type Graphics struct {
	frameTarget     uint64
	frameTargetTime float64
}

func makeGraphics() Graphics {
	return Graphics{
		frameTarget:     60,
		frameTargetTime: 1000.0 / 60.0,
	}
}

func (e *Graphics) FrameTarget(v uint64) {
	e.frameTarget = v
	e.frameTargetTime = 1000.0 / float64(v)
}
