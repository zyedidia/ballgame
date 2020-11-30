package main

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var errBegin = errors.New("begin")

const (
	mw = 160
	mh = 160

	diffHard       = 0
	diffImpossible = 1
)

type Menu struct {
	active     bool
	tick       float64
	bg         *Map
	difficulty int
}

func NewMenu(md *Map) *Menu {
	return &Menu{
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
		return errBegin
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		m.difficulty++
		m.difficulty %= 2
	}
	m.tick += 0.05

	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	m.bg.Draw(screen)

	ebitenutil.DebugPrintAt(screen, "Controls: Left/Right to move", width/8, 14)
	ebitenutil.DebugPrintAt(screen, "          Up to jump", width/8, 34)
	ebitenutil.DebugPrintAt(screen, "          Shift to dash", width/8, 54)
	ebitenutil.DebugPrintAt(screen, "          Escape to quit", width/8, 74)
	ebitenutil.DebugPrintAt(screen, "Difficulty: "+m.DifficultyString(), width/8, 94)
	ebitenutil.DebugPrintAt(screen, "Press enter to begin", width/4, 124)
}

func (m *Menu) DifficultyString() string {
	switch m.difficulty {
	case diffHard:
		return "hard"
	case diffImpossible:
		return "impossible"
	default:
		return "error: unknown difficulty"
	}
}

func (m *Menu) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}
