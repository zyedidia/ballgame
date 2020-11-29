package main

import (
	"image"
	_ "image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/shurcooL/httpfs/vfsutil"
)

var assets *AssetManager

type AssetManager struct {
	dir    string
	images map[string]*ebiten.Image
	music  map[string]*audio.Player
	sound  map[string]*audio.Player
	anims  map[string]*Animation
}

func NewAssets(dir string) *AssetManager {
	return &AssetManager{
		dir:    dir,
		images: make(map[string]*ebiten.Image),
		music:  make(map[string]*audio.Player),
		sound:  make(map[string]*audio.Player),
		anims:  make(map[string]*Animation),
	}
}

func (a *AssetManager) LoadSound() error {
	dir := filepath.Join(a.dir, "sound")

	return vfsutil.Walk(assetfs, dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ".ogg" {
			return nil
		}

		f, err := assetfs.Open(path)
		if err != nil {
			return err
		}
		s, err := vorbis.Decode(audioctx, f)
		if err != nil {
			return err
		}
		player, err := audio.NewPlayer(audioctx, s)
		if err != nil {
			return err
		}
		a.sound[info.Name()] = player
		log.Println("Loaded", path)

		return nil
		// return f.Close()
	})
}

func (a *AssetManager) LoadMusic() error {
	dir := filepath.Join(a.dir, "music")

	return vfsutil.Walk(assetfs, dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ".mp3" {
			return nil
		}

		f, err := assetfs.Open(path)
		if err != nil {
			return err
		}
		s, err := mp3.Decode(audioctx, f)
		if err != nil {
			return err
		}
		player, err := audio.NewPlayer(audioctx, s)
		if err != nil {
			return err
		}
		a.music[info.Name()] = player
		log.Println("Loaded", path)

		return nil
	})
}

func (a *AssetManager) LoadImages() error {
	dir := filepath.Join(a.dir, "img")

	return vfsutil.Walk(assetfs, dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ".png" {
			return nil
		}

		f, err := assetfs.Open(path)
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
	return a.images[name]
}

func (a *AssetManager) GetSound(name string) *audio.Player {
	return a.sound[name]
}

func (a *AssetManager) GetMusic(name string) *audio.Player {
	return a.music[name]
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
