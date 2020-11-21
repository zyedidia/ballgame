package main

import (
	"image"

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
				tiles[i][j].Kind = 1
			}
			tiles[i][0].Kind = 1
			tiles[i][j].W = 16
			tiles[i][j].H = 16
			tiles[i][j].X = j * 16
			tiles[i][j].Y = i * 16
		}
	}

	return &MapData{
		Bgs:       []string{"bg_0.png", "bg_1.png", "bg_2.png"},
		Tilesheet: "tileset.png",
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

func LoadMap(space *Space, data *MapData) *Map {
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

func (m *Map) Update() error {
	return nil
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
