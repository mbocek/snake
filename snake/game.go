package snake

import (
	"github.com/asaskevich/EventBus"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rs/zerolog/log"
	"os"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

var (
	bus = EventBus.New()
)

type Game struct {
	layout   *Layout
	input    *Input
	gameOver bool
}

func NewGame(sizeX, sizeY int) *Game {
	g := &Game{
		input: NewInput(),
	}
	g.layout = NewLayout(sizeX, sizeY)
	if err := bus.Subscribe(topicSnakeOutOfBoard, g.gameOverHandler); err != nil {
		log.Fatal().Err(err).Str("topic", topicSnakeOutOfBoard).Msg("cannot subscribe to topic")
	}
	if err := bus.Subscribe(topicRestart, g.restartHandler); err != nil {
		log.Fatal().Err(err).Str("topic", topicRestart).Msg("cannot subscribe to topic")
	}
	if err := bus.Subscribe(topicShutDown, g.shutDownHandler); err != nil {
		log.Fatal().Err(err).Str("topic", topicShutDown).Msg("cannot subscribe to topic")
	}
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	g.input.Update()
	g.layout.Update()
	return nil
}

func (g *Game) gameOverHandler(gameOver bool) {
	g.gameOver = gameOver
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameOver {
		ebitenutil.DebugPrint(screen, "Game Over!")
	} else {
		g.layout.Draw(screen)
	}
}

func (g *Game) restartHandler() {
	log.Debug().Msg("restart game")
	g.gameOver = false
}

func (g *Game) shutDownHandler() {
	log.Debug().Msg("shutdown game")
	os.Exit(0)
}
