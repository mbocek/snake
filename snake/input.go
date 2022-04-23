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

const (
	topicMove    = "topic:move"
	topicRestart = "topic:restart"
)

type Input struct {
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		bus.Publish(topicMove, topDirection)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		bus.Publish(topicMove, bottomDirection)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		bus.Publish(topicMove, rightDirection)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		bus.Publish(topicMove, leftDirection)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		bus.Publish(topicRestart)
	}
}
