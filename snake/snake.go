package snake

import (
	"errors"
	"math/rand"

	log "github.com/sirupsen/logrus"
)

const (
	topicSnakeOutOfBoard = "topic:snake:outOfBoard"
	topicFood            = "topic:snake:food"
)

type Snake struct {
	positions   Positions
	tiles       Tiles
	boardWidth  int
	boardHeight int
}

type Position struct {
	x, y int
}

type Positions []Position

var (
	ErrorOutOfBoard = errors.New("Snake is out of theBoard")
)

func NewSnake(tiles Tiles) *Snake {
	s := &Snake{
		positions: make(Positions, 0),
		tiles:     tiles,
	}
	s.boardWidth = len(s.tiles[0])
	s.boardHeight = len(s.tiles)
	s.init()
	return s
}

func (s *Snake) moveUp() {
	position := s.positions[0]
	newPosition := Position{x: position.x, y: position.y - 1}
	s.update(s.positions, newPosition)
}

func (s *Snake) moveDown() {
	position := s.positions[0]
	newPosition := Position{x: position.x, y: position.y + 1}
	s.update(s.positions, newPosition)
}

func (s *Snake) moveLeft() {
	position := s.positions[0]
	newPosition := Position{x: position.x - 1, y: position.y}
	s.update(s.positions, newPosition)
}

func (s *Snake) moveRight() {
	position := s.positions[0]
	newPosition := Position{x: position.x + 1, y: position.y}
	log.Debug("Position: %w", newPosition)
	s.update(s.positions, newPosition)
}

func (s *Snake) Move(direction int) {
	switch direction {
	case topDirection:
		s.moveUp()
	case bottomDirection:
		s.moveDown()
	case rightDirection:
		s.moveRight()
	case leftDirection:
		s.moveLeft()
	}
}

func (s *Snake) update(positions Positions, newPosition Position) {
	if !s.checkOutOfBoard(newPosition) {
		s.positions = append(Positions{newPosition}, s.positions...)
		if s.tiles[newPosition.y][newPosition.x].IsFood() {
			s.tiles[newPosition.y][newPosition.x].Empty()
			bus.Publish(topicFood, 1)
		} else {
			s.positions = s.positions[:len(s.positions)-1]
		}
	}
}

func (s *Snake) checkOutOfBoard(position Position) bool {
	if position.x < 0 || position.x-1 > s.boardWidth || position.y < 0 || position.y-1 > s.boardHeight {
		bus.Publish(topicSnakeOutOfBoard, true)
		return true
	}
	return false
}

func (s *Snake) init() {
	s.positions = make(Positions, 0)
	s.positions = append(s.positions, Position{x: rand.Intn(s.boardWidth), y: rand.Intn(s.boardHeight)})
}
