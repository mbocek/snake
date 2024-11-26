package snake

import (
	"math/rand"
)

const (
	topicSnakeOutOfBoard = "topic:snake:outOfBoard"
	topicFood            = "topic:snake:food"
	topicIncreaseSpeed   = "topic:snake:increaseSpeed"
	lengthToIncrease     = 3
	speedIncrease        = 2
)

type Snake struct {
	positions   Positions
	tiles       *Tiles
	boardWidth  int
	boardHeight int
}

type Position struct {
	x, y int
}

type Positions []Position

func NewSnake(tiles *Tiles) *Snake {
	s := &Snake{
		positions: make(Positions, 0),
		tiles:     tiles,
	}
	s.boardWidth = len((*s.tiles)[0])
	s.boardHeight = len(*s.tiles)
	s.init()
	return s
}

func (s *Snake) moveUp() {
	position := s.positions[0]
	newPosition := Position{x: position.x, y: position.y - 1}
	s.update(newPosition)
}

func (s *Snake) moveDown() {
	position := s.positions[0]
	newPosition := Position{x: position.x, y: position.y + 1}
	s.update(newPosition)
}

func (s *Snake) moveLeft() {
	position := s.positions[0]
	newPosition := Position{x: position.x - 1, y: position.y}
	s.update(newPosition)
}

func (s *Snake) moveRight() {
	position := s.positions[0]
	newPosition := Position{x: position.x + 1, y: position.y}
	s.update(newPosition)
}

func (s *Snake) Move(direction Direction) {
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

func (s *Snake) Contains(x, y int) bool {
	for _, p := range s.positions {
		if p.x == x && p.y == y {
			return true
		}
	}
	return false
}

func (s *Snake) update(newPosition Position) {
	if !s.checkOutOfBoard(newPosition) {
		s.positions = append(Positions{newPosition}, s.positions...)
		tiles := *s.tiles
		if tiles[newPosition.y][newPosition.x].IsFood() {
			tiles[newPosition.y][newPosition.x].Empty()
			bus.Publish(topicFood, 1)
			if len(s.positions)%lengthToIncrease == 0 {
				bus.Publish(topicIncreaseSpeed, speedIncrease)
			}
		} else {
			s.positions = s.positions[:len(s.positions)-1]
		}
	}
}

func (s *Snake) checkOutOfBoard(position Position) bool {
	if position.x < 0 || position.x > s.boardWidth-1 || position.y < 0 || position.y > s.boardHeight-1 {
		bus.Publish(topicSnakeOutOfBoard, true)
		return true
	}
	return false
}

func (s *Snake) init() {
	s.positions = make(Positions, 0)
	s.positions = append(s.positions, Position{x: rand.Intn(s.boardWidth), y: rand.Intn(s.boardHeight)})
}
