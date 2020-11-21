package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
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

func NewTile(space *Space, tilesheet *ebiten.Image, rect image.Rectangle, x, y int) *Tile {
	shape := NewRectangleShape(cp.NewStaticBody(), "wall", x+rect.Dx()/2, y+rect.Dy()/2, rect.Dx(), rect.Dy())
	space.Add(shape)
	return &Tile{
		tilesheet: tilesheet,
		rect:      rect,
		shape:     shape,
	}
}

func (t *Tile) Update() error {
	return nil
}

func (t *Tile) Draw(screen *ebiten.Image) {
	pos := t.shape.Body().Position()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(t.rect.Dx()/2), -float64(t.rect.Dy()/2))
	op.GeoM.Translate(pos.X, pos.Y)

	screen.DrawImage(t.tilesheet.SubImage(t.rect).(*ebiten.Image), op)
}
