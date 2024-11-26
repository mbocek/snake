package snake

type TileType int

const (
	empty TileType = iota
	food
	snake
)

type Tile struct {
	tileType TileType
	repaint  bool
}

func NewTile(tileType TileType) *Tile {
	t := &Tile{
		tileType: tileType,
		repaint:  true,
	}
	return t
}

func (t *Tile) Snake() {
	if t.tileType != snake {
		t.tileType = snake
		t.repaint = true
	}
}

func (t *Tile) Empty() {
	if t.tileType != empty {
		t.tileType = empty
		t.repaint = true
	}
}

func (t *Tile) Type() TileType {
	return t.tileType
}

func (t *Tile) IsFood() bool {
	return t.tileType == food
}

func (t *Tile) IsSnake() bool {
	return t.tileType == snake
}

func (t *Tile) Repaint() bool {
	if t.repaint {
		t.repaint = false
		return true
	}
	return false
}
