package main

import (
	"image"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

type TileData struct {
	Kind int
	X, Y int
	W, H int
}

type Tile struct {
	tilesheet *ebiten.Image
	rect      image.Rectangle
	shape     *Shape
}

func NewTile(space *resolv.Space, tilesheet *ebiten.Image, rect image.Rectangle, x, y int) *Tile {
	shape := NewRectangleShape("wall", x, y, rect.Dx(), rect.Dy())
	shape.AddTo(space)
	return &Tile{
		tilesheet: tilesheet,
		rect:      rect,
		shape:     shape,
	}
}

func (t *Tile) Update(*resolv.Space) {
}

func (t *Tile) Draw(screen *ebiten.Image) {
	pos := t.shape.pos

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(t.rect.Dx()/2), -float64(t.rect.Dy()/2))
	op.GeoM.Translate(pos.X, pos.Y)

	screen.DrawImage(t.tilesheet.SubImage(t.rect).(*ebiten.Image), op)
}
