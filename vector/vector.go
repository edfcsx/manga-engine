package vector

import "golang.org/x/exp/constraints"

type Vec2[T constraints.Ordered] struct {
	X T
	Y T
}

func MakeVec2[T constraints.Ordered](x T, y T) Vec2[T] {
	return Vec2[T]{X: x, Y: y}
}

func Add[T constraints.Integer | constraints.Float](a Vec2[T], b Vec2[T]) Vec2[T] {
	return Vec2[T]{X: a.X + b.X, Y: a.Y + b.Y}
}

func Mul[T constraints.Integer | constraints.Float](a Vec2[T], b T) Vec2[T] {
	return Vec2[T]{X: a.X * b, Y: a.Y * b}
}

func MulVec2[T constraints.Integer | constraints.Float](a Vec2[T], b Vec2[T]) Vec2[T] {
	return Vec2[T]{X: a.X * b.X, Y: a.Y * b.Y}
}
