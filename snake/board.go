package snake

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tileSize = 10
)

type Board struct {
	sizeX      int
	sizeY      int
	tiles      Tiles
	tileImages TileImages
	xTiles     int
	yTiles     int
}

type Tiles [][]*Tile
type TileImages [][]*ebiten.Image

func NewBoard(sizeX, sizeY int) *Board {
	b := &Board{
		sizeX:  sizeX,
		sizeY:  sizeY,
		xTiles: sizeX / tileSize,
		yTiles: sizeY / tileSize,
	}
	return b
}

func (b *Board) Size() (int, int) {
	return b.sizeX, b.sizeY
}

func (b *Board) Update() error {
	return nil
}

func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(color.RGBA{0xbb, 0x00, 0x00, 0xaa})
	if !b.isTilesInitialized() {
		b.init()
	}
	b.drawTiles(boardImage)
}

func (b *Board) drawTiles(boardImage *ebiten.Image) {
	for i := range b.tiles {
		for j := range b.tiles[i] {
			b.tiles[i][j].Draw(b.tileImages[i][j])
			opTile := &ebiten.DrawImageOptions{}
			opTile.GeoM.Translate(float64(j*tileSize), float64(i*tileSize))
			boardImage.DrawImage(b.tileImages[i][j], opTile)
		}
	}
}

func (b *Board) init() {
	b.tiles = make(Tiles, b.yTiles)
	b.tileImages = make(TileImages, b.yTiles)
	for i := range b.tiles {
		b.tiles[i] = make([]*Tile, b.xTiles)
		b.tileImages[i] = make([]*ebiten.Image, b.xTiles)
		for j := range b.tiles[i] {
			b.tiles[i][j] = NewTile(TileType(b.randomTileType(100)))
			b.tileImages[i][j] = ebiten.NewImage(tileSize, tileSize)
		}
	}
}

func (b *Board) randomTileType(entropy int) int {
	return rand.Intn(entropy+1) / entropy
}

func (b *Board) isTilesInitialized() bool {
	return b.tiles != nil
}
