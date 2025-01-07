package manga

import (
	"github.com/veandco/go-sdl2/sdl"
	"manga_engine/vector"
	"math"
)

const (
	ShapeCircle = iota
	ShapeRectangle
	ShapePoint
	ShapeLine
)

type ColliderShape interface {
	GetType() int32
	X() int32
	Y() int32
	MoveTo(x int32, y int32)
	Render(transform *TransformComponent)
}

// PointShape ----------------- point shape -----------------
type PointShape struct {
	position vector.Vec2[int32]
}

func MakePointShape() *PointShape {
	return &PointShape{
		position: vector.MakeVec2[int32](0, 0),
	}
}

func (p *PointShape) GetType() int32 {
	return ShapePoint
}

func (p *PointShape) X() int32 {
	return p.position.X
}

func (p *PointShape) Y() int32 {
	return p.position.Y
}

func (p *PointShape) MoveTo(x int32, y int32) {
	p.position.X = x
	p.position.Y = y
}

func (p *PointShape) Render(transform *TransformComponent) {
	err := Engine.renderer.SetDrawColor(255, 0, 0, 255)

	if err != nil {
		return
	}

	err = Engine.renderer.DrawPoint(p.position.X, p.position.Y)

	if err != nil {
		return
	}
}

// ----------------- end point shape -----------------

// CircleShape ----------------- circle shape -----------------
type CircleShape struct {
	position vector.Vec2[int32]
	radius   int32
}

func MakeCircleShape(radius int32) *CircleShape {
	return &CircleShape{
		position: vector.MakeVec2[int32](0, 0),
		radius:   radius,
	}
}

func (c *CircleShape) X() int32 {
	return c.position.X
}

func (c *CircleShape) Y() int32 {
	return c.position.Y
}

func (c *CircleShape) GetType() int32 {
	return ShapeCircle
}

func (c *CircleShape) Radius() int32 {
	return c.radius
}

func (c *CircleShape) MoveTo(x int32, y int32) {
	c.position.X = x
	c.position.Y = y
}

func (c *CircleShape) Render(transform *TransformComponent) {
	err := Engine.renderer.SetDrawColor(255, 0, 0, 255)

	if err != nil {
		return
	}

	centerX := c.position.X + ((transform.size.X * transform.scale) / 2)
	centerY := c.position.Y + ((transform.size.Y * transform.scale) / 2)

	for angle := 0.0; angle < 360.0; angle += 1.0 {
		rad := angle * (math.Pi / 180.0)
		x := centerX + int32(float64(c.radius)*math.Cos(rad))
		y := centerY + int32(float64(c.radius)*math.Sin(rad))
		err = Engine.renderer.DrawPoint(x, y)
	}
}

// ----------------- end circle shape -----------------

// RectangleShape ----------------- rectangle shape -----------------
type RectangleShape struct {
	position vector.Vec2[int32]
	left     int32
	top      int32
	right    int32
	bottom   int32
}

func MakeRectangleShape(w int32, h int32) *RectangleShape {
	return &RectangleShape{
		position: vector.MakeVec2[int32](0, 0),
		left:     0,
		top:      0,
		right:    w,
		bottom:   h,
	}
}

func (r *RectangleShape) X() int32 {
	return r.position.X
}

func (r *RectangleShape) Y() int32 {
	return r.position.Y
}

func (r *RectangleShape) GetType() int32 {
	return ShapeRectangle
}

func (r *RectangleShape) Left() int32 {
	return r.position.X + r.left
}

func (r *RectangleShape) Top() int32 {
	return r.position.Y + r.top
}

func (r *RectangleShape) Right() int32 {
	return r.position.X + r.right
}

func (r *RectangleShape) Bottom() int32 {
	return r.position.Y + r.bottom
}

func (r *RectangleShape) MoveTo(x int32, y int32) {
	r.position.X = x
	r.position.Y = y
}

func (r *RectangleShape) Render(transform *TransformComponent) {
	err := Engine.renderer.SetDrawColor(255, 0, 0, 255)

	if err != nil {
		return
	}

	err = Engine.renderer.DrawRect(&sdl.Rect{
		X: r.Left(),
		Y: r.Top(),
		W: r.right - r.left,
		H: r.bottom - r.top,
	})
}

// ----------------- end rectangle shape -----------------

// LineShape ----------------- line shape -----------------
type LineShape struct {
	position vector.Vec2[int32]
	a        vector.Vec2[int32]
	b        vector.Vec2[int32]
}

func MakeLineShape(x1 int32, y1 int32, x2 int32, y2 int32) *LineShape {
	return &LineShape{
		position: vector.MakeVec2[int32](0, 0),
		a:        vector.MakeVec2[int32](x1, y1),
		b:        vector.MakeVec2[int32](x2, y2),
	}
}

func (l *LineShape) X() int32 {
	return l.position.X
}

func (l *LineShape) Y() int32 {
	return l.position.Y
}

func (l *LineShape) GetType() int32 {
	return ShapeLine
}

func (l *LineShape) AX() int32 {
	return l.position.X + l.a.X
}

func (l *LineShape) AY() int32 {
	return l.position.Y + l.a.Y
}

func (l *LineShape) BX() int32 {
	return l.position.X + l.b.X
}

func (l *LineShape) BY() int32 {
	return l.position.Y + l.b.Y
}

func (l *LineShape) MoveTo(px int32, py int32) {
	l.position.X = px
	l.position.Y = py
}

func (l *LineShape) Render(transform *TransformComponent) {
	err := Engine.renderer.SetDrawColor(255, 0, 0, 255)

	if err != nil {
		return
	}

	err = Engine.renderer.DrawLine(l.AX(), l.AY(), l.BX(), l.BY())
}

// ----------------- end line shape -----------------
