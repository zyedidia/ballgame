package main

import (
	"image"
	"math/rand"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

type MapData struct {
	Tilesheet string
	Bgs       []string
	Tiles     [][]TileData
	W, H      int
}

func DefaultMapData() *MapData {
	w, h := 15, 10

	tiles := make([][]TileData, h)
	for i := range tiles {
		tiles[i] = make([]TileData, w)
		for j := range tiles[i] {
			if i == 0 || i == len(tiles)-1 {
				tiles[i][j].Kind = rand.Intn(4) + 1
			}
			tiles[i][j].W = 16
			tiles[i][j].H = 16
			tiles[i][j].X = j*16 + 8
			tiles[i][j].Y = i*16 + 8
		}
		tiles[i][0].Kind = rand.Intn(4) + 1
		tiles[i][len(tiles[i])-1].Kind = rand.Intn(4) + 1
	}

	return &MapData{
		Bgs:       []string{"bg_0.png", "temple.png"},
		Tilesheet: "tiles.png",
		W:         w,
		H:         h,
		Tiles:     tiles,
	}
}

type Map struct {
	tilesheet *ebiten.Image
	bgs       []*ebiten.Image
	tiles     []*Tile
}

func LoadMap(space *resolv.Space, data *MapData) *Map {
	tilesheet := assets.GetImage(data.Tilesheet)
	bgs := make([]*ebiten.Image, len(data.Bgs))
	for i, bg := range data.Bgs {
		bgs[i] = assets.GetImage(bg)
	}

	tiles := make([]*Tile, 0, data.H*data.W/2)
	for _, row := range data.Tiles {
		for _, t := range row {
			var rect image.Rectangle
			switch t.Kind {
			case 1:
				rect = image.Rect(0, 0, t.W, t.H)
			case 2:
				rect = image.Rect(16, 0, 16+t.W, t.H)
			case 3:
				rect = image.Rect(0, 16, t.W, 16+t.H)
			case 4:
				rect = image.Rect(16, 16, 16+t.W, 16+t.H)
			default:
				continue
			}
			tiles = append(tiles, NewTile(space, tilesheet, rect, t.X, t.Y))
		}
	}

	return &Map{
		tilesheet: tilesheet,
		bgs:       bgs,
		tiles:     tiles,
	}
}

func (m *Map) Update(*resolv.Space) {
}

func (m *Map) Draw(screen *ebiten.Image) {
	for _, bg := range m.bgs {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(bg, op)
	}

	for _, t := range m.tiles {
		t.Draw(screen)
	}
}
