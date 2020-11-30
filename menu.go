package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	mw = 160
	mh = 160
)

type Menu struct {
	img    *ebiten.Image
	active bool
	tick   float64
	bg     *Map
}

func NewMenu(md *Map) *Menu {
	return &Menu{
		img:    assets.GetImage("calendar-stone.png"),
		active: true,
		tick:   0,
		bg:     md,
	}
}

func (m *Menu) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errGameEsc
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		m.active = false
	}
	m.tick += 0.05

	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	m.bg.Draw(screen)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(mw/2), -float64(mh/2))
	op.GeoM.Scale(0.02*math.Cos(m.tick)+0.85, 0.02*math.Cos(m.tick)+0.85)
	op.GeoM.Translate(width/2, height/2-10)
	// screen.DrawImage(m.img, op)
	ebitenutil.DebugPrintAt(screen, "Controls: Left/Right to move", width/8, 32)
	ebitenutil.DebugPrintAt(screen, "          Up to jump", width/8, 52)
	ebitenutil.DebugPrintAt(screen, "          Shift to dash", width/8, 72)
	ebitenutil.DebugPrintAt(screen, "          Escape to quit", width/8, 92)
	ebitenutil.DebugPrintAt(screen, "Press enter to begin", width/4, height-40)
}

func (m *Menu) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}
