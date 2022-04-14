package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	topDirection = iota
	bottomDirection
	rightDirection
	leftDirection
)

type Input struct {
	direction int
	restart   bool
}

func NewInput() *Input {
	return &Input{
		direction: topDirection,
	}
}

func (i *Input) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		i.direction = topDirection
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		i.direction = bottomDirection
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		i.direction = rightDirection
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		i.direction = leftDirection
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		i.restart = true
	}
}

func (i *Input) ClearRestart() bool {
	if i.restart {
		i.restart = false
	}
	return i.restart
}
