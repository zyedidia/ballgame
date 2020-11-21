package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

const (
	pw = 16
	ph = 16

	gravity = 0.1
)

type Player struct {
	input    Input
	shape    *Shape
	img      *AnimImage
	velocity Vec2f
}

func NewPlayer(space *Space, input Input, x, y int) *Player {
	shape := NewRectangleShape(cp.NewKinematicBody(), "player", x, y, pw, ph)
	space.Add(shape)
	return &Player{
		input: input,
		shape: shape,
		img: &AnimImage{
			anim:  assets.GetAnimation("player_idle"),
			count: 0,
		},
		velocity: Vec2f{0, 0},
	}
}

func (p *Player) SetAnimation(s string) {
	p.img.anim = assets.GetAnimation(s)
	p.img.count = 0
}

func (p *Player) Update(space *Space) error {
	pos := p.shape.Body().Position()
	p.velocity.X = p.input.Get(ActionRight) - p.input.Get(ActionLeft)

	dx, dy := p.velocity.X, p.velocity.Y
	p.velocity.Y += gravity

	fmt.Println(dx)

	walls := space.FilterByTags("wall")
	collision := walls.Resolve(p.shape.collider, 0, int32(dy))
	if collision.Colliding() {
		// pos.X += float64(collision.ResolveX)
		pos.Y += float64(collision.ResolveY)
		// p.velocity.X = 0
		p.velocity.Y = 0
	} else {
		if int(dy) != 0 {
			pos.Y += dy
		}
	}
	if int(dx) != 0 {
		pos.X += dx
	}

	p.shape.Body().SetPosition(pos)
	p.shape.Update()
	p.img.Update()
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	pos := p.shape.Body().Position()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-pw/2, -ph/2)
	op.GeoM.Translate(pos.X, pos.Y)

	screen.DrawImage(p.img.Image(), op)
}
