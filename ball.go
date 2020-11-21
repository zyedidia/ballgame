package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

const radius = 8

type Ball struct {
	shape *Shape
	img   *ebiten.Image
}

func NewBall(space *Space, x, y int) *Ball {
	mass := 10.0
	body := cp.NewBody(mass, cp.MomentForCircle(mass, 0, float64(radius), cp.Vector{}))
	body.SetVelocity(-100.0, 0.0)
	shape := NewCircleShape(body, "ball", x, y, radius)
	space.Add(shape)

	return &Ball{
		shape: shape,
		img:   assets.GetImage("ball_red.png"),
	}
}

func (b *Ball) Update(*Space) error {
	b.shape.Update()
	return nil
}

func (b *Ball) Draw(screen *ebiten.Image) {
	pos := b.shape.Body().Position()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-radius, -radius)
	op.GeoM.Rotate(b.shape.Body().Angle())
	op.GeoM.Translate(pos.X, pos.Y)

	screen.DrawImage(b.img, op)
}
