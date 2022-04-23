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
	emptyColor = color.RGBA{0x11, 0x11, 0x11, 0xff}
	foodColor  = color.RGBA{0x00, 0x99, 0x00, 0xff}
	snakeColor = color.RGBA{0xff, 0xff, 0x00, 0xff}
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

func (t *Tile) Snake() {
	t.tileType = snake
}

func (t *Tile) Clone() *Tile {
	return &Tile{
		tileType: t.tileType,
	}
}

func (t *Tile) IsFood() bool {
	return t.tileType == food
}

func (t *Tile) Empty() {
	t.tileType = empty
}
