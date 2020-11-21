package main

import (
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

type Space struct {
	*resolv.Space
	cpspace *cp.Space
}

func NewSpace() *Space {
	cpspace := cp.NewSpace()
	cpspace.Iterations = 10
	cpspace.SetGravity(cp.Vector{X: 0, Y: 500})
	cpspace.SleepTimeThreshold = 0.5

	return &Space{
		Space:   resolv.NewSpace(),
		cpspace: cpspace,
	}
}

func (s *Space) Add(shape *Shape) {
	s.Space.Add(shape.collider)
	s.cpspace.AddBody(shape.Body())
	s.cpspace.AddShape(shape.Shape)
}

func (s *Space) Step(dt float64) {
	s.cpspace.Step(1.0 / float64(ebiten.MaxTPS()))
}
