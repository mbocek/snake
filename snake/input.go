package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rs/zerolog/log"
)

type Direction int

const (
	noDirection     Direction = -1
	topDirection    Direction = 0
	bottomDirection Direction = 1
	rightDirection  Direction = 2
	leftDirection   Direction = 3
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
		log.Debug().Msg("press UP")
		bus.Publish(topicMove, topDirection)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		log.Debug().Msg("press DOWN")
		bus.Publish(topicMove, bottomDirection)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		log.Debug().Msg("press RIGHT")
		bus.Publish(topicMove, rightDirection)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		log.Debug().Msg("press LEFT")
		bus.Publish(topicMove, leftDirection)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		bus.Publish(topicRestart)
	}
}
