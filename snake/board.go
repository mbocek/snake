package snake

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"

	log "github.com/sirupsen/logrus"
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
	snakeDirection int
}

type Tiles [][]*Tile
type TileImages [][]*ebiten.Image

func NewBoard(sizeX, sizeY int) *Board {
	b := &Board{
		sizeX:          sizeX,
		sizeY:          sizeY,
		xTiles:         sizeX / tileSize,
		yTiles:         sizeY / tileSize,
		snakeDirection: -1,
	}
	b.init()
	bus.Subscribe(topicMove, b.directionHandler)
	bus.Subscribe(topicRestart, b.restartHandler)
	return b
}

func (b *Board) Size() (int, int) {
	return b.sizeX, b.sizeY
}

func (b *Board) Update() error {
	b.applySnake()
	b.snake.Move(b.snakeDirection)
	return nil
}

func (b *Board) Draw(boardImage *ebiten.Image) {
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
			b.tileImages[i][j] = ebiten.NewImage(tileSize, tileSize)
		}
		copy(b.tilesWithSnake[i], b.tiles[i])
	}
	b.snake = NewSnake(b.tiles)
}

func (b *Board) randomTileType(entropy int) int {
	return rand.Intn(entropy+1) / entropy
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

func (b *Board) directionHandler(direction int) {
	b.snakeDirection = direction
}

func (b *Board) restartHandler() {
	log.Debug("Restart board")
	b.snakeDirection = -1
	b.init()
}
