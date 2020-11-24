package main

import "github.com/hajimehoshi/ebiten/v2/audio"

const sampleRate = 44100

var audioctx *audio.Context

func init() {
	audioctx = audio.NewContext(sampleRate)
}
