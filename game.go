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
	space   *resolv.Space
	objects []GameObject
}

func NewGame(md *MapData) *Game {
	space := resolv.NewSpace()

	assets.GetMusic("music.mp3").Play()

	objects := make([]GameObject, 0, 10)
	objects = append(objects, LoadMap(space, md))
	NewSegmentShape("wall", 0, height, width, height).AddTo(space)
	NewSegmentShape("wall", 0, height, 0, 0).AddTo(space)
	NewSegmentShape("wall", 0, 0, width, 0).AddTo(space)
	NewSegmentShape("wall", width, height, width, 0).AddTo(space)

	objects = append(objects, NewBall(space, width/2, height/2))
	objects = append(objects, NewPlayer(space, NewKeyboard(KeyboardDefaults), width/2, height/2))
	objects = append(objects, NewRing(space, width/2, height/3))

	return &Game{
		space:   space,
		objects: objects,
	}
}

func (g *Game) Update() error {
	for _, o := range g.objects {
		o.Update(g.space)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errGameEsc
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, o := range g.objects {
		o.Draw(screen)
	}
	ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(score)+", Best: "+strconv.Itoa(best))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}
