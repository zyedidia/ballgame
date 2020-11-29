package main

import (
	"fmt"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	bw       = 8
	bh       = 8
	maxbally = 6

	bgravity = 0.3
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

	player := space.FilterByTags("player")
	if res := player.Resolve(b.shape.collider, 0, dy); res.Colliding() {
		b.shape.pos.Y += float64(res.ResolveY)
		if p, ok := res.ShapeB.GetData().(*Player); ok {
			b.velocity.Y -= float64(int(p.velocity.Y)) / 4
			fmt.Println(p.velocity.Y)
		}
		b.velocity.Y *= -1.1
	} else {
		walls := space.FilterByTags("wall")
		if res := walls.Resolve(b.shape.collider, dx, 0); res.Colliding() && !res.Teleporting {
			b.shape.pos.X += float64(res.ResolveX)
			b.velocity.X *= -1
			// s := assets.GetSound("bounce.ogg")
			// s.Rewind()
			// s.Play()
		} else {
			b.shape.pos.X += float64(dx)
		}

		if res := walls.Resolve(b.shape.collider, 0, dy); res.Colliding() && !res.Teleporting {
			b.shape.pos.Y += float64(res.ResolveY)
			b.velocity.Y *= -0.95
			// s := assets.GetSound("bounce.ogg")
			// s.Rewind()
			// s.Play()
		} else {
			b.shape.pos.Y += float64(dy)
		}
	}

	if b.velocity.Y > maxbally {
		b.velocity.Y = maxbally
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
