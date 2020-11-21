package main

import (
	"errors"
	"image"
	_ "image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

var assets *AssetManager

type AssetManager struct {
	dir    string
	images map[string]*ebiten.Image
	anims  map[string]*Animation
}

func NewAssets(dir string) *AssetManager {
	return &AssetManager{
		dir:    dir,
		images: make(map[string]*ebiten.Image),
		anims:  make(map[string]*Animation),
	}
}

func (a *AssetManager) LoadSound() error {
	return errors.New("Not implemented")
}

func (a *AssetManager) LoadMusic() error {
	return errors.New("Not implemented")
}

func (a *AssetManager) LoadImages() error {
	imagedir := filepath.Join(a.dir, "img")

	return filepath.Walk(imagedir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ".png" {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		img, _, err := image.Decode(f)
		if err != nil {
			return err
		}
		a.images[info.Name()] = ebiten.NewImageFromImage(img)
		log.Println("Loaded", path)

		return f.Close()
	})
}

func (a *AssetManager) LoadAnimations() error {
	a.anims = map[string]*Animation{
		"player_idle": &Animation{
			loop:     true,
			slowdown: 8,
			img:      a.GetImage("herochar_spritesheet.png"),
			frames:   buildFrames(0, 4, 16, 16, 4),
		},
		"player_run": &Animation{
			loop:     true,
			slowdown: 4,
			img:      a.GetImage("herochar_spritesheet.png"),
			frames:   buildFrames(0, 1, 16, 16, 6),
		},
	}
	return nil
}

func (a *AssetManager) GetImage(name string) *ebiten.Image {
	if img, ok := a.images[name]; ok {
		return img
	}
	log.Fatal("Could not load", name)
	return nil
}

func (a *AssetManager) GetAnimation(name string) *Animation {
	return a.anims[name]
}

func buildFrames(x, y, w, h, nframes int) []image.Rectangle {
	frames := make([]image.Rectangle, nframes)
	for i := 0; i < nframes; i++ {
		frames[i] = image.Rect(x*w+w*i, y*h, x*w+w*i+w, y*h+h)
	}
	return frames
}
