package main

import (
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject interface {
	Update(space *resolv.Space)
	Draw(screen *ebiten.Image)
}

type Shape struct {
	pos      Vec2f
	collider resolv.Shape
}

func NewCircleShape(tag string, x, y, r int) *Shape {
	collider := resolv.NewCircle(int32(x), int32(y), int32(r))
	collider.AddTags(tag)

	return &Shape{
		pos:      Vec2f{X: float64(x), Y: float64(y)},
		collider: collider,
	}
}

func NewSegmentShape(tag string, x, y, x2, y2 int) *Shape {
	collider := resolv.NewLine(int32(x), int32(y), int32(x2), int32(y2))
	collider.AddTags(tag)

	return &Shape{
		collider: collider,
	}
}

func NewRectangleShape(tag string, x, y, w, h int) *Shape {
	collider := resolv.NewRectangle(int32(x-w/2), int32(y-h/2), int32(w), int32(h))
	collider.AddTags(tag)

	return &Shape{
		pos:      Vec2f{X: float64(x), Y: float64(y)},
		collider: collider,
	}
}

func (s *Shape) Update() {
	x, y := s.pos.X, s.pos.Y
	switch t := s.collider.(type) {
	case *resolv.Rectangle:
		s.collider.SetXY(int32(x)-t.W/2, int32(y)-t.H/2)
	default:
		s.collider.SetXY(int32(x), int32(y))
	}
}

func (s *Shape) Pos() (float64, float64) {
	return s.pos.X, s.pos.Y
}

func (s *Shape) AddTo(space *resolv.Space) {
	space.Add(s.collider)
}
