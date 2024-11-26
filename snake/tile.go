package snake

type TileType int

const (
	TileEmpty TileType = iota
	TileFood
	TileSnake
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
	if t.tileType != TileSnake {
		t.tileType = TileSnake
		t.repaint = true
	}
}

func (t *Tile) Empty() {
	if t.tileType != TileEmpty {
		t.tileType = TileEmpty
		t.repaint = true
	}
}

func (t *Tile) Food() {
	if t.tileType != TileFood {
		t.tileType = TileFood
		t.repaint = true
	}
}

func (t *Tile) Type() TileType {
	return t.tileType
}

func (t *Tile) IsFood() bool {
	return t.tileType == TileFood
}

func (t *Tile) IsSnake() bool {
	return t.tileType == TileSnake
}

func (t *Tile) Repaint() bool {
	if t.repaint {
		t.repaint = false
		return true
	}
	return false
}
