package colliderCircle

import (
	"github.com/veandco/go-sdl2/sdl"
	colliderI "manga_engine/manga/interfaces/collider"
	componentI "manga_engine/manga/interfaces/component"
)

type Circle struct {
	x      int32
	y      int32
	radius int32
}

func Make(radius int32) *Circle {
	return &Circle{
		x:      0,
		y:      0,
		radius: radius,
	}
}

func (s *Circle) GetType() int32 {
	return colliderI.CircleType
}

func (s *Circle) X() int32 {
	return s.x
}

func (s *Circle) Y() int32 {
	return s.y
}

func (s *Circle) MoveTo(x, y int32) {
	s.x = x
	s.y = y
}

func (s *Circle) Render(t *componentI.TransformComponent, r *sdl.Renderer) {

}

func (s *Circle) CollidesWith(shape colliderI.Shape) bool {
	return false
}

func (s *Circle) GetRadius() int32 {
	return s.radius
}
