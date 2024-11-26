package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
	"image/color"
	"math/rand"
)

type Board struct {
	sizeX          int
	sizeY          int
	tileSize       int
	tiles          Tiles
	snake          *Snake
	emptyTileImage *ebiten.Image
	foodTileImage  *ebiten.Image
	snakeTileImage *ebiten.Image
	snakeDirection Direction
	speed          int
	speedStep      int
}

type Tiles [][]*Tile

func NewBoard(sizeX, sizeY int) *Board {
	emptyColor := color.RGBA{0x11, 0x11, 0x11, 0xff}
	foodColor := color.RGBA{0x00, 0x99, 0x00, 0xff}
	snakeColor := color.RGBA{0xff, 0xff, 0x00, 0xff}

	b := Board{sizeX: sizeX, sizeY: sizeY, tileSize: 10, speed: 10}
	b.emptyTileImage = ebiten.NewImage(b.tileSize, b.tileSize)
	b.emptyTileImage.Fill(emptyColor)
	b.foodTileImage = ebiten.NewImage(b.tileSize, b.tileSize)
	b.foodTileImage.Fill(foodColor)
	b.snakeTileImage = ebiten.NewImage(b.tileSize, b.tileSize)
	b.snakeTileImage.Fill(snakeColor)

	b.init()
	bus.Subscribe(topicMove, b.directionHandler)
	bus.Subscribe(topicRestart, b.restartHandler)

	return &b
}

func (b *Board) getNumberOfXTiles() int {
	return b.sizeX / b.tileSize
}

func (b *Board) getNumberOfYTiles() int {
	return b.sizeY / b.tileSize
}

func (b *Board) init() {
	b.tiles = make(Tiles, b.getNumberOfYTiles())
	for i := range b.tiles {
		b.tiles[i] = make([]*Tile, b.getNumberOfXTiles())
		for j := range b.tiles[i] {
			b.tiles[i][j] = NewTile(TileType(b.randomTileType(100)))
		}
	}
	b.snake = NewSnake(&b.tiles)
}

func (b *Board) randomTileType(entropy int) int {
	return rand.Intn(entropy+1) / entropy
}

func (b *Board) Update() error {
	if ebiten.TPS()/b.speed > b.speedStep {
		b.speedStep++
		return nil
	}
	b.speedStep = 0

	b.snake.Move(b.snakeDirection)

	for i := range b.tiles {
		for j := range b.tiles[i] {
			tile := b.tiles[i][j]
			if b.snake.Contains(j, i) {
				tile.Snake()
			} else {
				if tile.IsSnake() {
					tile.Empty()
				}
			}
		}
	}
	return nil
}

func (b *Board) Draw(boardImage *ebiten.Image) {
	for i := range b.tiles {
		for j := range b.tiles[i] {
			tile := b.tiles[i][j]
			if tile.Repaint() {
				opTile := &ebiten.DrawImageOptions{}
				opTile.GeoM.Translate(float64(j*b.tileSize), float64(i*b.tileSize))
				switch tile.Type() {
				case empty:
					boardImage.DrawImage(b.emptyTileImage, opTile)
				case food:
					boardImage.DrawImage(b.foodTileImage, opTile)
				case snake:
					boardImage.DrawImage(b.snakeTileImage, opTile)
				}
			}
		}
	}
}

func (b *Board) Size() (int, int) {
	return b.sizeX, b.sizeY
}

func (b *Board) directionHandler(direction Direction) {
	b.snakeDirection = direction
}

func (b *Board) restartHandler() {
	log.Debug().Msg("restart board")
	b.snakeDirection = -1
	b.init()
}
