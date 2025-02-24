package manga

import (
	"fmt"
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

func (p *PointShape) Distance(x int32, y int32) float64 {
	deltaX := math.Abs(float64(x - p.position.X))
	deltaY := math.Abs(float64(y - p.position.Y))

	return math.Sqrt((deltaX * deltaX) + (deltaY * deltaY))
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

// ----------------- Collisions Resolvers -----------

func CollisionPointPoint(a *PointShape, b *PointShape) bool {
	return a.X() == b.X() && a.Y() == b.Y()
}

func CollisionPointRect(p *PointShape, r *RectangleShape) bool {
	if p.X() >= r.Left() && p.X() <= r.Right() {
		if p.Y() >= r.Top() && p.Y() <= r.Bottom() {
			return true
		}
	}

	return false
}

func CollisionPointCircle(p *PointShape, c *CircleShape) bool {
	return p.Distance(c.X(), c.Y()) <= float64(c.Radius())
}

func CollisionPointLine(p *PointShape, l *LineShape) bool {
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

func CollisionRectRect(a *RectangleShape, b *RectangleShape) bool {
	// verifica se existe sobreposição nos dois eixos
	overlapX := b.Left() <= a.Right() && a.Left() <= b.Right()
	overlapY := b.Top() <= a.Bottom() && a.Top() <= b.Bottom()

	return overlapX && overlapY
}

func CollisionRectCircle(r *RectangleShape, c *CircleShape) bool {
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

	// verifica se o ponto mais próximo está dentro do círculo
	return CollisionPointCircle(&PointShape{position: vector.MakeVec2[int32](px, py)}, c)
}

func CollisionResolver(a, b ColliderShape) bool {
	collision := false

	switch a.GetType() {
	case ShapePoint:
		switch b.GetType() {
		case ShapePoint:
			collision = CollisionPointPoint(a.(*PointShape), b.(*PointShape))
			break
		case ShapeRectangle:
			collision = CollisionPointRect(a.(*PointShape), b.(*RectangleShape))
			break
		case ShapeCircle:
			collision = CollisionPointCircle(a.(*PointShape), b.(*CircleShape))
			break
		case ShapeLine:
			collision = CollisionPointLine(a.(*PointShape), b.(*LineShape))
		}
	case ShapeRectangle:
		switch b.GetType() {
		case ShapePoint:
			collision = CollisionRect
		}
	}

	return collision
}

// ----------------- end Collision Resolvers --------

const ColliderSystemID string = "ColliderSystem"

type ColliderType int32

const (
	ColliderStatic ColliderType = iota
	ColliderMoving
)

type Collider struct {
	shape       ColliderShape
	onCollision func(*Entity)
}

// ColliderSystem TODO: change to use hashmap for faster lookup
type ColliderSystem struct {
	static []*Collider
	moving []*Collider
}

func MakeColliderSystem() *ColliderSystem {
	return &ColliderSystem{}
}

func (c *ColliderSystem) Initialize() {}

func (c *ColliderSystem) Update() {
	if len(c.moving) == 0 || (len(c.moving) == 1) && (len(c.static) == 0) {
		return
	}

	// compare all move objects first
	for idx, moving := range c.moving {
		for idx2 := idx + 1; idx2 < len(c.moving); idx2++ {
			if CollisionResolver(moving.shape, c.moving[idx2].shape) {
				moving.onCollision(c.moving[idx2].shape.(*ColliderComponent).Entity)
				c.moving[idx2].onCollision(moving.shape.(*ColliderComponent).Entity)
			}
		}
	}

	// compare all move objects with static objects
	for _, moving := range c.moving {
		for _, static := range c.static {

		}
	}

	fmt.Println("ColliderSystem Update")
}

func (c *ColliderSystem) Register(t ColliderType, shape ColliderShape, onCollision func(*Entity)) {
	if t == ColliderStatic {
		c.static = append(c.static, &Collider{shape: shape, onCollision: onCollision})
	} else {
		c.moving = append(c.moving, &Collider{shape: shape, onCollision: onCollision})
	}
}
