package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
	boardSize    = 4
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
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	g.input.Update()
	if err := g.layout.Update(g.input); err != nil {
		g.gameOver = true
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameOver {
		ebitenutil.DebugPrint(screen, "Game Over!")
	} else {
		g.layout.Draw(screen)
	}
}
