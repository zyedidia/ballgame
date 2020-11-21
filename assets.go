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
}

func NewAssets(dir string) *AssetManager {
	return &AssetManager{
		dir:    dir,
		images: make(map[string]*ebiten.Image),
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

func (a *AssetManager) GetImage(name string) *ebiten.Image {
	if img, ok := a.images[name]; ok {
		return img
	}
	log.Fatal("Could not load", name)
	return nil
}
