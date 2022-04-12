package snake

import "math/rand"

type Snake struct {
	positions Positions
}

type Position struct {
	x, y int
}

type Positions []Position

func NewSnake(boardWidth, boardHeight int) *Snake {
	s := &Snake{
		positions: make(Positions, 0),
	}
	s.positions = append(s.positions, Position{x: rand.Intn(boardWidth), y: rand.Intn(boardHeight)})
	return s
}

func (s *Snake) MoveUp(tiles Tiles) {
	position := s.positions[0]
	newPosition := Position{x: position.x, y: position.y - 1}
	s.update(tiles, s.positions, newPosition)
}

func (s *Snake) MoveDown(tiles Tiles) {
	position := s.positions[0]
	newPosition := Position{x: position.x, y: position.y + 1}
	s.update(tiles, s.positions, newPosition)
}

func (s *Snake) MoveLeft(tiles Tiles) {
	position := s.positions[0]
	newPosition := Position{x: position.x - 1, y: position.y}
	s.update(tiles, s.positions, newPosition)
}

func (s *Snake) MoveRight(tiles Tiles) {
	position := s.positions[0]
	newPosition := Position{x: position.x + 1, y: position.y}
	s.update(tiles, s.positions, newPosition)
}
func (s *Snake) update(tiles Tiles, positions Positions, newPosition Position) {
	s.positions = append(Positions{newPosition}, s.positions...)
	if tiles[newPosition.y][newPosition.x].IsFood() {
		tiles[newPosition.y][newPosition.x].Empty()
	} else {
		s.positions = s.positions[:len(s.positions)-1]
	}
}

func (s *Snake) Append(x, y int) {
	s.positions = append(s.positions, Position{x: x, y: y})
}
