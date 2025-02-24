package colliderI

import (
	"github.com/veandco/go-sdl2/sdl"
	componentI "manga_engine/manga/interfaces/component"
)

type Shape interface {
	X() int32
	Y() int32
	MoveTo(x, y int32)
	GetType() int32
	CollidesWith(shape Shape) bool
	Render(t componentI.TransformComponent, r *sdl.Renderer)
}

const (
	PointType = iota
	LineType
	CircleType
	RectangleType
)

type PointShape interface {
	Shape
	Distance(x, y int32) float64
}

type CircleShape interface {
	Shape
	GetRadius() float64
}

type LineShape interface {
	Shape
	AX() int32
	AY() int32
	BX() int32
	BY() int32
}

type RectangleShape interface {
	Shape
	Left() int32
	Right() int32
	Top() int32
	Bottom() int32
}
