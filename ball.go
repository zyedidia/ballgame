package main

import (
	"time"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	bw       = 8
	bh       = 8
	maxbally = 6
	maxballx = 2

	minjitter = -1.5
	maxjitter = 1.5

	bgravity = 0.3
)

type Ball struct {
	shape    *Shape
	img      *ebiten.Image
	velocity Vec2f
	paused   *AtomicBool
}

func NewBall(space *resolv.Space, x, y int) *Ball {
	shape := NewRectangleShape("ball", x, y, bw, bh)
	shape.AddTo(space)

	b := &Ball{
		shape:    shape,
		img:      assets.GetImage("ball.png"),
		velocity: Vec2f{1, 0},
		paused:   NewBool(true),
	}

	go func() {
		b.Reset()
	}()
	shape.collider.SetData(b)
	return b
}

func (b *Ball) Reset() {
	b.shape.pos.X = width / 2
	b.shape.pos.Y = height / 2
	b.velocity = Vec2f{0, 0}
	b.paused.SetTo(true)
	go func() {
		<-time.After(1 * time.Second)
		b.paused.SetTo(false)
	}()
}

func (b *Ball) Update(space *resolv.Space) {
	if b.paused.IsSet() {
		return
	}

	b.velocity.Y += bgravity

	dx, dy := int32(b.velocity.X), int32(b.velocity.Y)

	player := space.FilterByTags("player")
	if res := player.Resolve(b.shape.collider, 0, dy); res.Colliding() {
		b.shape.pos.Y += float64(res.ResolveY)
		if p, ok := res.ShapeB.GetData().(*Player); ok {
			if b.shape.pos.Y <= p.shape.pos.Y {
				b.velocity.Y -= float64(int(p.velocity.Y)) / 4
				b.velocity.Y *= -1.1
				if b.shape.pos.X > p.shape.pos.X {
					b.velocity.X += 1
				} else if b.shape.pos.X < p.shape.pos.X {
					b.velocity.X -= 1
				}
				if b.velocity.X > maxballx {
					b.velocity.X = maxballx
				}
				if b.velocity.X < -maxballx {
					b.velocity.X = -maxballx
				}
			}
		}
	} else {
		walls := space.FilterByTags("wall")
		if res := walls.Resolve(b.shape.collider, dx, 0); res.Colliding() && !res.Teleporting {
			b.shape.pos.X += float64(res.ResolveX)
			b.velocity.X *= -1
		} else {
			b.shape.pos.X += float64(dx)
		}

		if res := walls.Resolve(b.shape.collider, 0, dy); res.Colliding() && !res.Teleporting {
			b.shape.pos.Y += float64(res.ResolveY)
			b.velocity.Y *= -1
			if dy > 0 {
				score = 0
				b.Reset()
			}
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
