package colliderPoint

import (
	"github.com/veandco/go-sdl2/sdl"
	"manga_engine/manga/colors"
	colliderI "manga_engine/manga/interfaces/collider"
	componentI "manga_engine/manga/interfaces/component"
	"math"
)

type Point struct {
	x int32
	y int32
}

func Make() *Point {
	return &Point{
		x: 0,
		y: 0,
	}
}

func (p *Point) GetType() int32 {
	return colliderI.PointType
}

func (p *Point) X() int32 {
	return p.x
}

func (p *Point) Y() int32 {
	return p.y
}

func (p *Point) MoveTo(x, y int32) {
	p.x = x
	p.y = y
}

func (p *Point) Render(t *componentI.TransformComponent, r *sdl.Renderer) {
	// TODO: add errors in log
	c := colors.RED
	err := r.SetDrawColor(c.R, c.G, c.B, c.A)
	if err != nil {
		return
	}

	err = r.DrawPoint(p.x, p.y)
	if err != nil {
		return
	}
}

func (p *Point) Distance(x int32, y int32) float64 {
	deltaX := math.Abs(float64(x - p.x))
	deltaY := math.Abs(float64(y - p.y))

	return math.Sqrt((deltaX * deltaX) + (deltaY * deltaY))
}

func (p *Point) CollidesWith(shape colliderI.Shape) bool {
	switch shape.GetType() {
	case colliderI.PointType:
		return collidesPoint(p, shape)
	case colliderI.CircleType:
		return collidesCircle(p, shape)
	case colliderI.RectangleType:
		return collidesRectangle(p, shape)
	case colliderI.LineType:
		return collidesLine(p, shape)
	default:
		return false
	}
}

func collidesPoint(p colliderI.PointShape, b colliderI.Shape) bool {
	return p.X() == b.X() && p.Y() == b.X()
}

func collidesCircle(p colliderI.PointShape, b interface{}) bool {
	circle, ok := b.(colliderI.CircleShape)

	if !ok {
		//TODO: remove panic and log error
		panic(ok)
	}

	return p.Distance(circle.X(), circle.Y()) <= circle.GetRadius()
}

func collidesRectangle(p colliderI.PointShape, b interface{}) bool {
	rect, ok := b.(colliderI.RectangleShape)

	if !ok {
		panic(ok)
	}

	if p.X() >= rect.Left() && p.X() <= rect.Right() {
		if p.Y() >= rect.Top() && p.Y() <= rect.Bottom() {
			return true
		}
	}

	return false
}

func collidesLine(p colliderI.PointShape, b interface{}) bool {
	l, ok := b.(colliderI.LineShape)

	if !ok {
		panic(ok)
	}

	// Coordenadas do ponto
	px, py := float64(p.X()), float64(p.Y())

	// Coordenadas dos pontos da linha
	ax, ay := float64(l.AX()), float64(l.AY())
	bx, by := float64(l.BX()), float64(l.BY())

	// Calcular a distância perpendicular do ponto à linha
	numerator := math.Abs((by-ay)*px - (bx-ax)*py + bx*ay - by*ax)
	denominator := math.Sqrt((by-ay)*(by-ay) + (bx-ax)*(bx-ax))
	distance := numerator / denominator

	// Verificar se o ponto está dentro dos limites da linha
	if distance == 0 {
		if (px >= math.Min(ax, bx) && px <= math.Max(ax, bx)) && (py >= math.Min(ay, by) && py <= math.Max(ay, by)) {
			return true
		}
	}

	return false
}
