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
	speed          *Speed
	foodCount      int
}

type Tiles [][]*Tile

func NewBoard(sizeX, sizeY int, speed *Speed) *Board {
	emptyColor := color.RGBA{0x11, 0x11, 0x11, 0xff}
	foodColor := color.RGBA{0x00, 0x99, 0x00, 0xff}
	snakeColor := color.RGBA{0xff, 0xff, 0x00, 0xff}

	const tileSize = 10
	const foodCount = 20
	b := Board{sizeX: sizeX, sizeY: sizeY, tileSize: tileSize, speed: speed, snakeDirection: noDirection, foodCount: foodCount}
	b.emptyTileImage = ebiten.NewImage(b.tileSize, b.tileSize)
	b.emptyTileImage.Fill(emptyColor)
	b.foodTileImage = ebiten.NewImage(b.tileSize, b.tileSize)
	b.foodTileImage.Fill(foodColor)
	b.snakeTileImage = ebiten.NewImage(b.tileSize, b.tileSize)
	b.snakeTileImage.Fill(snakeColor)

	b.init()
	if err := bus.Subscribe(topicMove, b.directionHandler); err != nil {
		log.Fatal().Err(err).Str("topic", topicMove).Msg("cannot subscribe to topic")
	}
	if err := bus.Subscribe(topicRestart, b.restartHandler); err != nil {
		log.Fatal().Err(err).Str("topic", topicRestart).Msg("cannot subscribe to topic")
	}
	if err := bus.Subscribe(topicIncreaseSpeed, b.speed.Increase); err != nil {
		log.Fatal().Err(err).Str("topic", topicIncreaseSpeed).Msg("cannot subscribe to topic")
	}
	if err := bus.Subscribe(topicFood, b.addFood); err != nil {
		log.Fatal().Err(err).Str("topic", topicFood).Msg("cannot subscribe to topic")
	}

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
			b.tiles[i][j] = NewTile(TileEmpty)
		}
	}
	b.addFood(b.foodCount)
	b.snake = NewSnake(&b.tiles)
}

func (b *Board) randomTileType(entropy int) int {
	return rand.Intn(entropy+1) / entropy
}

func (b *Board) Update() error {
	if !b.speed.Over() {
		return nil
	}
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
				case TileEmpty:
					boardImage.DrawImage(b.emptyTileImage, opTile)
				case TileFood:
					boardImage.DrawImage(b.foodTileImage, opTile)
				case TileSnake:
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

func (b *Board) addFood(count int) {
	for range count {
		for {
			x := rand.Intn(b.getNumberOfXTiles())
			y := rand.Intn(b.getNumberOfYTiles())
			if b.tiles[y][x].Type() == TileEmpty {
				b.tiles[y][x].Food()
				break
			}
		}
	}
}
