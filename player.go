package main

import (
	"fmt"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	pw = 16
	ph = 16
	cw = 11
	ch = 16

	maxspd      = 2.0
	friction    = 0.2
	accel       = 0.5
	jumpspd     = 6.0
	force_accel = 1.1

	pgravity = 0.4
)

type Player struct {
	input    Input
	shape    *Shape
	img      *AnimImage
	velocity Vec2f
	flip     bool
	force    float64
}

func NewPlayer(space *resolv.Space, input Input, x, y int) *Player {
	shape := NewRectangleShape("player", x, y, cw, ch)
	shape.AddTo(space)
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
}

func (p *Player) Update(space *resolv.Space) {
	hit := p.input.Get(ActionHit)
	jump := p.input.Get(ActionJump)
	move := p.input.Get(ActionRight) - p.input.Get(ActionLeft)

	if move != 0 {
		p.SetAnimation("player_run")
	} else {
		p.SetAnimation("player_idle")
	}

	// p.flip unchanged if move = 0
	if move < 0 {
		p.flip = true
	} else if move > 0 {
		p.flip = false
	}

	p.velocity.Y += pgravity

	if p.velocity.X > friction {
		p.velocity.X -= friction
	} else if p.velocity.X < -friction {
		p.velocity.X += friction
	} else {
		p.velocity.X = 0
	}

	p.velocity.X += accel * move

	if p.velocity.X > maxspd {
		p.velocity.X = maxspd
	} else if p.velocity.X < -maxspd {
		p.velocity.X = -maxspd
	}

	dx, dy := int32(p.velocity.X), int32(p.velocity.Y)

	balls := space.FilterByTags("ball")
	collision := balls.Resolve(p.shape.collider, dx, dy)
	if collision.Colliding() {
		if hit > 0 {
			ball := collision.ShapeB.GetData().(*Ball)
			fmt.Println("Yes")
			ball.velocity.Y = 10
		}
	}

	down := space.Resolve(p.shape.collider, 0, ph/2)
	onGround := down.Colliding()

	if jump > 0 && onGround {
		p.velocity.Y = -jumpspd
	}

	walls := space.FilterByTags("wall")
	if res := walls.Resolve(p.shape.collider, dx, 0); res.Colliding() {
		dx = res.ResolveX
		p.velocity.X = 0
	}
	p.shape.pos.X += float64(dx)

	if res := walls.Resolve(p.shape.collider, 0, dy); res.Colliding() {
		dy = res.ResolveY
		p.velocity.Y = 0
	}
	p.shape.pos.Y += float64(dy)

	p.shape.Update()
	p.img.Update()
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-pw/2, -ph/2)
	if p.flip {
		op.GeoM.Scale(-1, 1)
	}
	op.GeoM.Translate(p.shape.pos.X, p.shape.pos.Y)

	screen.DrawImage(p.img.Image(), op)
}
