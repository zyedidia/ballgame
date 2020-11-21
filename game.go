package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp"
)

type Game struct {
	space   *Space
	objects []GameObject
}

func NewGame(md *MapData) *Game {
	space := NewSpace()

	objects := make([]GameObject, 0, 10)
	objects = append(objects, LoadMap(space, md))
	space.Add(NewSegmentShape(cp.NewStaticBody(), "wall", 0, height, width, height))
	space.Add(NewSegmentShape(cp.NewStaticBody(), "wall", 0, height, 0, 0))
	space.Add(NewSegmentShape(cp.NewStaticBody(), "wall", 0, 0, width, 0))
	space.Add(NewSegmentShape(cp.NewStaticBody(), "wall", width, height, width, 0))

	objects = append(objects, NewBall(space, width/2, height/2))
	objects = append(objects, NewPlayer(space, NewKeyboard(KeyboardDefaults), width/2, height/2))

	return &Game{
		space:   space,
		objects: objects,
	}
}

func (g *Game) Update() error {
	// update all the objects, but if any one of them has an error then
	// return the error after the update
	var err error
	for _, o := range g.objects {
		e := o.Update(g.space)
		if e != nil {
			err = e
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errGameEsc
	}

	g.space.Step(1.0 / float64(ebiten.MaxTPS()))

	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, o := range g.objects {
		o.Draw(screen)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}
