package main

import (
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	bw = 8
	bh = 8

	bgravity = 0.4
)

type Ball struct {
	shape    *Shape
	img      *ebiten.Image
	velocity Vec2f
}

func NewBall(space *resolv.Space, x, y int) *Ball {
	shape := NewRectangleShape("ball", x, y, bw, bh)
	shape.AddTo(space)

	b := &Ball{
		shape:    shape,
		img:      assets.GetImage("ball.png"),
		velocity: Vec2f{1, 0},
	}
	shape.collider.SetData(b)
	return b
}

func (b *Ball) Update(space *resolv.Space) {
	b.velocity.Y += bgravity

	dx, dy := int32(b.velocity.X), int32(b.velocity.Y)
	walls := space.FilterByTags("wall")
	if res := walls.Resolve(b.shape.collider, dx, 0); res.Colliding() && !res.Teleporting {
		b.shape.pos.X += float64(res.ResolveX)
		b.velocity.X *= -1
	} else {
		b.shape.pos.X += float64(dx)
	}

	if res := walls.Resolve(b.shape.collider, 0, dy); res.Colliding() && !res.Teleporting {
		b.shape.pos.Y += float64(res.ResolveY)
		b.velocity.Y *= -0.95
	} else {
		b.shape.pos.Y += float64(dy)
	}

	b.shape.Update()
}

func (b *Ball) Draw(screen *ebiten.Image) {
	x, y := b.shape.Pos()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-bw/2, -bh/2)
	op.GeoM.Translate(x, y)

	screen.DrawImage(b.img, op)
}
