package snake

import "github.com/hajimehoshi/ebiten/v2"

type Speed struct {
	speed       int
	currentStep int
}

func NewSpeed(speed int) *Speed {
	return &Speed{
		speed: speed,
	}
}

func (s *Speed) Over() bool {
	s.currentStep++
	if s.currentStep >= ebiten.TPS()/s.speed {
		s.Reset()
		return true
	}
	return false
}

func (s *Speed) Reset() {
	s.currentStep = 0
}

func (s *Speed) Increase(increment int) {
	s.speed += increment
}

func (s *Speed) Speed() int {
	return s.speed
}
