package main

import (
	"time"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rw  = 9
	rh  = 14
	crw = 9
	crh = 14
)

type Ring struct {
	shape     *Shape
	img       *ebiten.Image
	lastpoint time.Time // ugly hack
}

func NewRing(space *resolv.Space, x, y int) *Ring {
	shape := NewRectangleShape("ring", x, y, crw, crh)
	shape.AddTo(space)

	r := &Ring{
		shape:     shape,
		img:       assets.GetImage("ring.png"),
		lastpoint: time.Now(),
	}

	return r
}

func (r *Ring) Update(space *resolv.Space) {
	ball := space.FilterByTags("ball")
	if res := ball.Resolve(r.shape.collider, 0, 1); res.Colliding() {
		if b, ok := res.ShapeB.GetData().(*Ball); ok {
			if b.velocity.X != 0 && time.Since(r.lastpoint) >= 500*time.Millisecond {
				r.lastpoint = time.Now()
				score++
				if score > best {
					best = score
				}
				sfx := assets.GetSound("score.ogg")
				sfx.Rewind()
				sfx.Play()
			}
		}
	}
}

func (r *Ring) Draw(screen *ebiten.Image) {
	pos := r.shape.pos

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(rw/2), -float64(rh/2))
	op.GeoM.Translate(pos.X, pos.Y)

	screen.DrawImage(r.img, op)
}
