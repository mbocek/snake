package snake

import (
	"github.com/asaskevich/EventBus"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	log "github.com/sirupsen/logrus"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
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
	bus.Subscribe(topicSnakeOutOfBoard, g.gameOverHandler)
	bus.Subscribe(topicRestart, g.restartHandler)
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
	log.Debug("Restart game")
	g.gameOver = false
}
