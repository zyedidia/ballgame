package main

import (
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	pw = 16
	ph = 16
	cw = 14
	ch = 16

	maxspd      = 2.0
	friction    = 0.2
	accel       = 0.4
	jumpspd     = 5.0
	dashspd     = 10.0
	dashlen     = 2
	force_accel = 1.1

	pgravity = 0.3
)

type Player struct {
	input     Input
	shape     *Shape
	img       *AnimImage
	velocity  Vec2f
	flip      bool
	force     float64
	dashcount int
	dashdir   int
}

func NewPlayer(space *resolv.Space, input Input, x, y int) *Player {
	shape := NewRectangleShape("player", x, y, cw, ch)
	shape.AddTo(space)
	p := &Player{
		input: input,
		shape: shape,
		img: &AnimImage{
			anim:  assets.GetAnimation("player_idle"),
			count: 0,
		},
		velocity: Vec2f{0, 0},
	}
	shape.collider.SetData(p)
	return p
}

func (p *Player) SetAnimation(s string) {
	p.img.anim = assets.GetAnimation(s)
}

func (p *Player) Update(space *resolv.Space) {
	// hit := p.input.Get(ActionHit)
	jump := p.input.Get(ActionJump)
	move := p.input.Get(ActionRight) - p.input.Get(ActionLeft)
	dash := p.input.GetJustPressed(ActionDash)

	if move != 0 {
		p.SetAnimation("player_run")
	} else {
		p.SetAnimation("player_idle")
	}

	if dash > 0 && move != 0 {
		p.dashcount = 1
		if move > 0 {
			p.dashdir = 1
		} else {
			p.dashdir = -1
		}
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

	if p.dashcount > 0 {
		p.velocity.X = float64(dashspd * p.dashdir)

		if p.dashcount >= dashlen {
			p.dashcount = 0
		} else {
			p.dashcount++
		}
	} else {
		p.velocity.X += accel * move

		if p.velocity.X > maxspd {
			p.velocity.X = maxspd
		} else if p.velocity.X < -maxspd {
			p.velocity.X = -maxspd
		}
	}

	dx, dy := int32(p.velocity.X), int32(p.velocity.Y)

	walls := space.FilterByTags("wall")
	down := walls.Resolve(p.shape.collider, 0, ph/2)
	onGround := down.Colliding()

	if jump > 0 && onGround {
		p.velocity.Y = -jumpspd
	}

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
