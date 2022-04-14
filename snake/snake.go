package snake

import (
	"errors"
	"math/rand"
)

type Snake struct {
	positions   Positions
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

func NewSnake(boardWidth, boardHeight int) *Snake {
	s := &Snake{
		positions:   make(Positions, 0),
		boardWidth:  boardWidth,
		boardHeight: boardHeight,
	}
	s.positions = append(s.positions, Position{x: rand.Intn(boardWidth), y: rand.Intn(boardHeight)})
	return s
}

func (s *Snake) MoveUp(tiles Tiles) error {
	position := s.positions[0]
	newPosition := Position{x: position.x, y: position.y - 1}
	return s.update(tiles, s.positions, newPosition)
}

func (s *Snake) MoveDown(tiles Tiles) error {
	position := s.positions[0]
	newPosition := Position{x: position.x, y: position.y + 1}
	return s.update(tiles, s.positions, newPosition)
}

func (s *Snake) MoveLeft(tiles Tiles) error {
	position := s.positions[0]
	newPosition := Position{x: position.x - 1, y: position.y}
	return s.update(tiles, s.positions, newPosition)
}

func (s *Snake) MoveRight(tiles Tiles) error {
	position := s.positions[0]
	newPosition := Position{x: position.x + 1, y: position.y}
	return s.update(tiles, s.positions, newPosition)
}

func (s *Snake) update(tiles Tiles, positions Positions, newPosition Position) error {
	if err := s.checkOutOfBoard(newPosition); err != nil {
		return err
	}

	s.positions = append(Positions{newPosition}, s.positions...)
	if tiles[newPosition.y][newPosition.x].IsFood() {
		tiles[newPosition.y][newPosition.x].Empty()
	} else {
		s.positions = s.positions[:len(s.positions)-1]
	}

	return nil
}

func (s *Snake) checkOutOfBoard(position Position) error {
	if position.x < 0 || position.x > s.boardWidth || position.y < 0 || position.y > s.boardHeight {
		return ErrorOutOfBoard
	}
	return nil
}
