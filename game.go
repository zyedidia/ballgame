package main

import (
	"strconv"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var score int
var best int

type Game struct {
	menu    *Menu
	space   *resolv.Space
	objects []GameObject
}

func NewGame(md *MapData) *Game {
	space := resolv.NewSpace()

	m := LoadMap(space, md)
	objects := make([]GameObject, 0, 10)
	objects = append(objects, LoadMap(space, md))

	return &Game{
		menu:    NewMenu(m),
		space:   space,
		objects: objects,
	}
}

func (g *Game) Init(difficulty int) {
	NewSegmentShape("wall", 0, height, width, height).AddTo(g.space)
	NewSegmentShape("wall", 0, height, 0, 0).AddTo(g.space)
	NewSegmentShape("wall", 0, 0, width, 0).AddTo(g.space)
	NewSegmentShape("wall", width, height, width, 0).AddTo(g.space)

	g.objects = append(g.objects, NewBall(g.space, width/2, height/2))
	g.objects = append(g.objects, NewPlayer(g.space, NewKeyboard(KeyboardDefaults), width/2, height/2))
	switch difficulty {
	case diffHard:
		g.objects = append(g.objects, NewRing(g.space, width/2, height/3))
	case diffImpossible:
		g.objects = append(g.objects, NewRing(g.space, width/2, height/8))
	}
}

func (g *Game) Update() error {
	if g.menu.active {
		r := g.menu.Update()
		if r == errBegin {
			g.Init(g.menu.difficulty)
		}
	}

	for _, o := range g.objects {
		o.Update(g.space)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errGameEsc
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.menu.active {
		g.menu.Draw(screen)
		return
	}

	for _, o := range g.objects {
		o.Draw(screen)
	}
	ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(score)+", Best: "+strconv.Itoa(best))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}
