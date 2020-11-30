package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var errGameEsc = errors.New("Escape pressed")

const (
	width  = 240
	height = 160
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Ball Game")
	ebiten.SetWindowResizable(true)

	assets = NewAssets("/")
	err := assets.LoadImages()
	if err != nil {
		log.Fatal(err)
	}
	err = assets.LoadMusic()
	if err != nil {
		log.Fatal(err)
	}
	err = assets.LoadSound()
	if err != nil {
		log.Fatal(err)
	}
	assets.LoadAnimations()

	assets.GetMusic("music.mp3").Play()

	g := NewGame(DefaultMapData())
	if err := ebiten.RunGame(g); err != nil {
		if err == errGameEsc {
			fmt.Println("Quit")
			os.Exit(0)
		}
		log.Fatal(err)
	}
}
