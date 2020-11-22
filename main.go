package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var errGameEsc = errors.New("Escape pressed")

const (
	width  = 240
	height = 160
)

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Ball Game")
	ebiten.SetWindowResizable(true)

	assets = NewAssets("assets")
	err := assets.LoadImages()
	if err != nil {
		log.Fatal(err)
	}
	assets.LoadAnimations()

	// bytes, _ := json.MarshalIndent(DefaultMapData(), "", "    ")
	// fmt.Println(string(bytes))

	g := NewGame(DefaultMapData())
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
