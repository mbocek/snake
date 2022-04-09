package snake

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileType int

const (
	empty TileType = iota
	food
	snake
)

var (
	emptyColor = color.RGBA{0x00, 0x00, 0x00, 0xff}
	foodColor  = color.RGBA{0x00, 0x99, 0x0, 0xff}
	snakeColor = color.RGBA{0x00, 0xff, 0x00, 0xff}
)

type Tile struct {
	tileType TileType
}

func NewTile(tileType TileType) *Tile {
	t := &Tile{
		tileType: tileType,
	}
	return t
}

func (t *Tile) Update() error {
	return nil
}

func (t *Tile) Draw(tileImage *ebiten.Image) {
	switch t.tileType {
	case empty:
		tileImage.Fill(emptyColor)
	case food:
		tileImage.Fill(foodColor)
	case snake:
		tileImage.Fill(snakeColor)
	}
}
