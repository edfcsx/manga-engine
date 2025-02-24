package colliderCircle

import (
	"github.com/veandco/go-sdl2/sdl"
	colliderPoint "manga_engine/manga/collider/point"
	"manga_engine/manga/colors"
	colliderI "manga_engine/manga/interfaces/collider"
	componentI "manga_engine/manga/interfaces/component"
	"math"
)

type Circle struct {
	x      int32
	y      int32
	radius float64
}

func Make(radius float64) *Circle {
	return &Circle{
		x:      0,
		y:      0,
		radius: radius,
	}
}

func (c *Circle) GetType() int32 {
	return colliderI.CircleType
}

func (c *Circle) X() int32 {
	return c.x
}

func (c *Circle) Y() int32 {
	return c.y
}

func (c *Circle) MoveTo(x, y int32) {
	c.x = x
	c.y = y
}

func (c *Circle) Render(t componentI.TransformComponent, renderer *sdl.Renderer) {
	color := colors.RED
	err := renderer.SetDrawColor(color.R, color.G, color.B, color.A)

	if err != nil {
		return
	}

	scale := t.GetScale()

	centerX := c.X() + ((t.GetSize().X * scale) / 2)
	centerY := c.Y() + ((t.GetSize().Y * scale) / 2)

	for angle := 0.0; angle < 360.0; angle += 1.0 {
		rad := angle * (math.Pi / 180.0)
		x := centerX + int32(c.radius*math.Cos(rad))
		y := centerY + int32(c.radius*math.Sin(rad))
		err = renderer.DrawPoint(x, y)
	}
}

func (c *Circle) GetRadius() float64 {
	return c.radius
}

func (c *Circle) CollidesWith(shape colliderI.Shape) bool {
	switch shape.GetType() {
	case colliderI.PointType:
		return collidesPoint(c, shape)
	case colliderI.CircleType:
		return collidesCircle(c, shape)
	case colliderI.RectangleType:
		return collidesRectangle(c, shape)
	case colliderI.LineType:
		return collidesLine(c, shape)

	default:
		return false
	}
}

func collidesPoint(c colliderI.CircleShape, b interface{}) bool {
	point, ok := b.(colliderI.PointShape)

	if !ok {
		panic(ok)
	}

	return point.CollidesWith(c)
}

func collidesCircle(c colliderI.CircleShape, b interface{}) bool {
	circle, ok := b.(colliderI.CircleShape)

	if !ok {
		panic(ok)
	}

	// deltas podem ser negativos se a subtração é feita na ordem errada
	// levando essa possibilidade em contra é melhor pegar os valores absolutos
	deltaX := math.Abs(float64(c.X() - circle.X()))
	deltaY := math.Abs(float64(c.Y() - circle.Y()))

	distance := math.Sqrt((deltaX * deltaX) + (deltaY * deltaY))

	if distance <= c.GetRadius()+circle.GetRadius() {
		return true
	}

	return false
}

func collidesRectangle(c colliderI.CircleShape, b interface{}) bool {
	r, ok := b.(colliderI.RectangleShape)

	if !ok {
		panic(ok)
	}

	// encontra o ponto do retângulo mais próximo do círculo
	var px, py int32

	if c.X() < r.Left() {
		px = r.Left()
	} else {
		if c.X() > r.Right() {
			px = r.Right()
		} else {
			px = c.X()
		}
	}

	if c.Y() < r.Top() {
		py = r.Top()
	} else {
		if c.Y() > r.Bottom() {
			py = r.Bottom()
		} else {
			py = c.Y()
		}
	}

	// verifica se existe colisão entre o ponto e o círculo
	point := colliderPoint.Make()
	point.MoveTo(px, py)

	return point.CollidesWith(c)
}

func collidesLine(c colliderI.CircleShape, b interface{}) bool {
	l, ok := b.(colliderI.LineShape)
	if !ok {
		panic(ok)
	}

	// Coordenadas do centro do círculo
	cx, cy := float64(c.X()), float64(c.Y())

	// Coordenadas dos pontos da linha
	ax, ay := float64(l.AX()), float64(l.AY())
	bx, by := float64(l.BX()), float64(l.BY())

	// Calcular a distância perpendicular do centro do círculo à linha
	numerator := math.Abs((by-ay)*cx - (bx-ax)*cy + bx*ay - by*ax)
	denominator := math.Sqrt((by-ay)*(by-ay) + (bx-ax)*(bx-ax))
	distance := numerator / denominator

	// Verificar se a distância é menor ou igual ao raio do círculo
	if distance <= c.GetRadius() {
		// Verificar se o ponto de interseção perpendicular está dentro dos limites da linha
		dotProduct := ((cx - ax) * (bx - ax)) + ((cy - ay) * (by - ay))
		lineLengthSquared := ((bx - ax) * (bx - ax)) + ((by - ay) * (by - ay))
		projection := dotProduct / lineLengthSquared

		if projection >= 0 && projection <= 1 {
			return true
		}
	}

	return false
}
