package snake

import "github.com/hajimehoshi/ebiten/v2"

type Layout struct {
	board      *Board
	info       *Info
	boardImage *ebiten.Image
	infoImage  *ebiten.Image
}

func NewLayout(layoutSizeX, layoutSizeY int) *Layout {
	l := &Layout{}
	l.board = NewBoard(layoutSizeX-infoWidth, layoutSizeY)
	l.info = NewInfo(infoWidth, layoutSizeY)

	return l
}

func (l *Layout) Update(input *Input) error {
	if err := l.board.Update(input); err != nil {
		return err
	}
	if err := l.info.Update(input); err != nil {
		return err
	}
	return nil
}

func (l *Layout) Draw(screen *ebiten.Image) {

	if l.boardImage == nil {
		w, h := l.board.Size()
		l.boardImage = ebiten.NewImage(w, h)
	}

	if l.infoImage == nil {
		w, h := l.info.Size()
		l.infoImage = ebiten.NewImage(w, h)
	}

	l.board.Draw(l.boardImage)
	opBoard := &ebiten.DrawImageOptions{}
	opBoard.GeoM.Translate(float64(0), float64(0))
	screen.DrawImage(l.boardImage, opBoard)

	l.info.Draw(l.infoImage)
	opInfo := &ebiten.DrawImageOptions{}
	opInfo.GeoM.Translate(float64(l.board.sizeX), float64(0))
	screen.DrawImage(l.infoImage, opInfo)
}
