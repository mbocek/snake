package snake

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
	boardSize    = 4
)

type Game struct {
	layout *Layout
	input  *Input
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
	return g.layout.Update(g.input)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xfa, 0xf8, 0xef, 0xff})
	g.layout.Draw(screen)
}
