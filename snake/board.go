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
	sizeX          int
	sizeY          int
	tiles          Tiles
	tilesWithSnake Tiles
	tileImages     TileImages
	xTiles         int
	yTiles         int
	snake          *Snake
	initDone       bool
}

type Tiles [][]*Tile
type TileImages [][]*ebiten.Image

func NewBoard(sizeX, sizeY int) *Board {
	b := &Board{
		sizeX:    sizeX,
		sizeY:    sizeY,
		xTiles:   sizeX / tileSize,
		yTiles:   sizeY / tileSize,
		initDone: false,
	}
	b.snake = NewSnake(b.xTiles, b.yTiles)
	return b
}

func (b *Board) Size() (int, int) {
	return b.sizeX, b.sizeY
}

func (b *Board) Update(input *Input) error {
	if b.initDone {
		b.move(input)
		b.applySnake()
	}
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
			b.tilesWithSnake[i][j].Draw(b.tileImages[i][j])
			opTile := &ebiten.DrawImageOptions{}
			opTile.GeoM.Translate(float64(j*tileSize), float64(i*tileSize))
			boardImage.DrawImage(b.tileImages[i][j], opTile)
		}
	}
}

func (b *Board) init() {
	b.tiles = make(Tiles, b.yTiles)
	b.tilesWithSnake = make(Tiles, b.yTiles)
	b.tileImages = make(TileImages, b.yTiles)
	for i := range b.tiles {
		b.tiles[i] = make([]*Tile, b.xTiles)
		b.tilesWithSnake[i] = make([]*Tile, b.xTiles)
		b.tileImages[i] = make([]*ebiten.Image, b.xTiles)
		for j := range b.tiles[i] {
			b.tiles[i][j] = NewTile(TileType(b.randomTileType(100)))
			b.tilesWithSnake[i][j] = b.tiles[i][j].Clone()
			b.tileImages[i][j] = ebiten.NewImage(tileSize, tileSize)
		}
	}
	b.initDone = true
}

func (b *Board) randomTileType(entropy int) int {
	return rand.Intn(entropy+1) / entropy
}

func (b *Board) isTilesInitialized() bool {
	return b.tiles != nil
}

func (b *Board) move(input *Input) {
	switch input.direction {
	case topDirection:
		b.snake.MoveUp(b.tiles)
	case bottomDirection:
		b.snake.MoveDown(b.tiles)
	case rightDirection:
		b.snake.MoveRight(b.tiles)
	case leftDirection:
		b.snake.MoveLeft(b.tiles)
	}
}

func (b *Board) applySnake() {
	for i, row := range b.tiles {
		for j := range row {
			b.tilesWithSnake[i][j] = b.tiles[i][j].Clone()
		}
	}
	for _, v := range b.snake.positions {
		b.tilesWithSnake[v.y][v.x].Snake()
	}
}
