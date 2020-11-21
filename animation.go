package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	img      *ebiten.Image
	frames   []image.Rectangle
	slowdown int
	loop     bool
}

type AnimImage struct {
	anim  *Animation
	count int
}

func (a *AnimImage) Update() {
	a.count++
}

func (a *AnimImage) Frame() int {
	frame := a.count / a.anim.slowdown
	if a.anim.loop {
		return frame % len(a.anim.frames)
	}
	return min(frame, len(a.anim.frames))
}

func (a *AnimImage) Image() *ebiten.Image {
	return a.anim.img.SubImage(a.anim.frames[a.Frame()]).(*ebiten.Image)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
