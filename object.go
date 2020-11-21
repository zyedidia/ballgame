package main

import (
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

type GameObject interface {
	Update(space *Space) error
	Draw(screen *ebiten.Image)
}

type Shape struct {
	*cp.Shape
	collider resolv.Shape
}

func NewCircleShape(b *cp.Body, tag string, x, y, r int) *Shape {
	b.SetPosition(cp.Vector{X: float64(x), Y: float64(y)})
	shape := cp.NewCircle(b, float64(r), cp.Vector{})
	shape.SetElasticity(0.8)
	shape.SetFriction(0.6)
	// TODO: mass, elasticity, friction

	collider := resolv.NewCircle(int32(x), int32(y), int32(r))
	collider.AddTags(tag)

	return &Shape{
		Shape:    shape,
		collider: collider,
	}
}

func NewSegmentShape(b *cp.Body, tag string, x, y, x2, y2 int) *Shape {
	shape := cp.NewSegment(b, cp.Vector{X: float64(x), Y: float64(y)}, cp.Vector{X: float64(x2), Y: float64(y2)}, 0)
	shape.SetElasticity(1)
	shape.SetFriction(1)
	collider := resolv.NewLine(int32(x), int32(y), int32(x2), int32(y2))
	collider.AddTags(tag)

	return &Shape{
		Shape:    shape,
		collider: collider,
	}
}

func NewRectangleShape(b *cp.Body, tag string, x, y, w, h int) *Shape {
	b.SetPosition(cp.Vector{X: float64(x), Y: float64(y)})
	shape := cp.NewBox(b, float64(w), float64(h), 0)
	shape.SetElasticity(1)
	shape.SetFriction(1)
	collider := resolv.NewRectangle(int32(x-w/2), int32(y-h/2), int32(w), int32(h))
	collider.AddTags(tag)

	return &Shape{
		Shape:    shape,
		collider: collider,
	}
}

// No need to call Update on static bodies
func (s *Shape) Update() {
	pos := s.Body().Position()
	switch t := s.collider.(type) {
	case *resolv.Rectangle:
		s.collider.SetXY(int32(pos.X)-t.W/2, int32(pos.Y)-t.H/2)
	default:
		s.collider.SetXY(int32(pos.X), int32(pos.Y))
	}
}
